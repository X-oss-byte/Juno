package pubsub_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/NethermindEth/juno/pubsub"
	"github.com/NethermindEth/juno/utils"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/require"
)

func testRegistry(t *testing.T) *pubsub.Registry {
	log := utils.NewNopZapLogger()
	registry := pubsub.New(log)
	t.Cleanup(func() {
		registry.Close()
	})
	return registry
}

func TestReliableSubscription(t *testing.T) {
	t.Parallel()

	subscribed := false
	t.Cleanup(func() {
		require.True(t, subscribed)
	})
	sub := event.NewSubscription(func(quit <-chan struct{}) error {
		subscribed = true
		<-quit
		return nil
	})

	registry := testRegistry(t)
	id := registry.Add(context.Background(), sub)
	err := registry.Delete(id)
	require.NoError(t, err)
}

func TestUnreliableSubscription(t *testing.T) {
	t.Parallel()

	sub := event.NewSubscription(func(_ <-chan struct{}) error {
		return errors.New("test err")
	})

	registry := testRegistry(t)
	id := registry.Add(context.Background(), sub)
	<-time.After(time.Millisecond * 100)
	err := registry.Delete(id)
	require.ErrorIs(t, err, pubsub.ErrNotFound, "should have unsubscribed already")
}

func TestEarlyShutdown(t *testing.T) {
	t.Parallel()

	sub := event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})

	registry := testRegistry(t)
	ctx, cancel := context.WithCancel(context.Background())
	id := registry.Add(ctx, sub)
	cancel()
	<-time.After(time.Millisecond * 100)
	err := registry.Delete(id)
	require.ErrorIs(t, err, pubsub.ErrNotFound, "should have unsubscribed already")
}
