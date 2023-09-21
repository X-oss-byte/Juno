package pubsub

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"sync"

	"github.com/NethermindEth/juno/utils"
	"github.com/ethereum/go-ethereum/event"
	"github.com/sourcegraph/conc"
)

var ErrNotFound = errors.New("subscription not found")

type Registry struct {
	log utils.SimpleLogger
	wg  conc.WaitGroup

	mu            sync.Mutex // protects subscriptions
	subscriptions map[uint64]event.Subscription
}

func New(log utils.SimpleLogger) *Registry {
	return &Registry{
		subscriptions: make(map[uint64]event.Subscription),
		log:           log,
	}
}

func getRandomID() uint64 {
	var n uint64
	for err := binary.Read(rand.Reader, binary.LittleEndian, &n); err != nil; {
	}
	return n
}

func (r *Registry) Add(ctx context.Context, sub event.Subscription) uint64 {
	id := getRandomID()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.subscriptions[id] = sub
	r.wg.Go(func() {
		select {
		case <-ctx.Done():
		case err := <-sub.Err():
			if err != nil {
				r.log.Warnw("Subscription failed", "err", err)
			}
		}
		// Unsubscribe releases the subscription resources.
		if err := r.Delete(id); err != nil {
			r.log.Warnw("Failed to unsubscribe", "subscriptionID", id, "err", err)
		}
	})
	return id
}

func (r *Registry) Delete(id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	sub, ok := r.subscriptions[id]
	if !ok {
		return ErrNotFound
	}
	sub.Unsubscribe()
	delete(r.subscriptions, id)
	return nil
}

// Close returns when all subscriptions have been handled.
func (r *Registry) Close() {
	r.wg.Wait()
}
