package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/NethermindEth/juno/node"
	"github.com/NethermindEth/juno/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const greeting = `
       _                    
      | |                   
      | |_   _ _ __   ___   
  _   | | | | | '_ \ / _ \  
 | |__| | |_| | | | | (_) |  
  \____/ \__,_|_| |_|\___/ %s

Juno is a Go implementation of a Starknet full-node client created by Nethermind.

`

const (
	configF              = "config"
	logLevelF            = "log-level"
	httpF                = "http"
	httpHostF            = "http-host"
	httpPortF            = "http-port"
	wsF                  = "ws"
	wsHostF              = "ws-host"
	wsPortF              = "ws-port"
	dbPathF              = "db-path"
	networkF             = "network"
	ethNodeF             = "eth-node"
	pprofF               = "pprof"
	pprofHostF           = "pprof-host"
	pprofPortF           = "pprof-port"
	colourF              = "colour"
	pendingPollIntervalF = "pending-poll-interval"
	p2pF                 = "p2p"
	p2pAddrF             = "p2p-addr"
	p2pBootPeersF        = "p2p-boot-peers"
	metricsF             = "metrics"
	metricsHostF         = "metrics-host"
	metricsPortF         = "metrics-port"
	grpcF                = "grpc"
	grpcHostF            = "grpc-host"
	grpcPortF            = "grpc-port"

	defaultConfig              = ""
	defaulHost                 = "localhost"
	defaultHTTP                = false
	defaultHTTPPort            = 6060
	defaultWS                  = false
	defaultWSPort              = 6061
	defaultEthNode             = ""
	defaultPprof               = false
	defaultPprofPort           = 6062
	defaultColour              = true
	defaultPendingPollInterval = time.Duration(0)
	defaultP2p                 = false
	defaultP2pAddr             = ""
	defaultP2pBootPeers        = ""
	defaultMetrics             = false
	defaultMetricsPort         = 9090
	defaultGRPC                = false
	defaultGRPCPort            = 6064

	configFlagUsage   = "The yaml configuration file."
	logLevelFlagUsage = "Options: debug, info, warn, error."
	httpUsage         = "Enables the HTTP RPC server on the default port and interface."
	httpHostUsage     = "The interface on which the HTTP RPC server will listen for requests."
	httpPortUsage     = "The port on which the HTTP server will listen for requests."
	wsUsage           = "Enables the Websocket RPC server on the default port."
	wsHostUsage       = "The interface on which the Websocket RPC server will listen for requests."
	wsPortUsage       = "The port on which the websocket server will listen for requests."
	dbPathUsage       = "Location of the database files."
	networkUsage      = "Options: mainnet, goerli, goerli2, integration."
	pprofUsage        = "Enables the pprof endpoint on the default port."
	pprofHostUsage    = "The interface on which the pprof HTTP server will listen for requests."
	pprofPortUsage    = "The port on which the pprof HTTP server will listen for requests."
	colourUsage       = "Uses --colour=false command to disable colourized outputs (ANSI Escape Codes)."
	ethNodeUsage      = "Websocket endpoint of the Ethereum node. In order to verify the correctness of the L2 chain, " +
		"Juno must connect to an Ethereum node and parse events in the Starknet contract."
	pendingPollIntervalUsage = "Sets how frequently pending block will be updated (disabled by default)"
	p2pUsage                 = "enable p2p server"
	p2PAddrUsage             = "specify p2p source address as multiaddr"
	p2pBootPeersUsage        = "specify list of p2p boot peers splitted by a comma"
	metricsUsage             = "Enables the prometheus metrics endpoint on the default port."
	metricsHostUsage         = "The interface on which the prometheus endpoint will listen for requests."
	metricsPortUsage         = "The port on which the prometheus endpoint will listen for requests."
	grpcUsage                = "Enable the HTTP GRPC server on the default port."
	grpcHostUsage            = "The interface on which the GRPC server will listen for requests."
	grpcPortUsage            = "The port on which the GRPC server will listen for requests."
)

var Version string

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		cancel()
	}()

	config := new(node.Config)
	cmd := NewCmd(config, func(cmd *cobra.Command, _ []string) error {
		fmt.Printf(greeting, Version)
		cfg := make(map[string]interface{})
		err := mapstructure.Decode(config, &cfg)
		if err != nil {
			return err
		}
		yamlConfig, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}
		fmt.Printf("Running Juno with Config:\n%s\n\n", string(yamlConfig))

		n, err := node.New(config, Version)
		if err != nil {
			return err
		}

		n.Run(cmd.Context())
		return nil
	})

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

// NewCmd returns a command that can be executed with any of the Cobra Execute* functions.
// The RunE field is set to the user-provided run function, allowing for robust testing setups.
//
//  1. NewCmd is called with a non-nil config and a run function.
//  2. An Execute* function is called on the command returned from step 1.
//  3. The config struct is populated.
//  4. Cobra calls the run function.
func NewCmd(config *node.Config, run func(*cobra.Command, []string) error) *cobra.Command { //nolint:funlen
	junoCmd := &cobra.Command{
		Use:     "juno [flags]",
		Short:   "Starknet client implementation in Go.",
		Version: Version,
		RunE:    run,
	}

	var cfgFile string

	// PreRunE populates the configuration struct from the Cobra flags and Viper configuration.
	// This is called in step 3 of the process described above.
	junoCmd.PreRunE = func(cmd *cobra.Command, _ []string) error {
		v := viper.New()
		if cfgFile != "" {
			v.SetConfigType("yaml")
			v.SetConfigFile(cfgFile)
			if err := v.ReadInConfig(); err != nil {
				return err
			}
		}

		if err := v.BindPFlags(cmd.Flags()); err != nil {
			return nil
		}

		// TextUnmarshallerHookFunc allows us to unmarshal values that satisfy the
		// encoding.TextUnmarshaller interface (see the LogLevel type for an example).
		err := v.Unmarshal(config, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
			mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToTimeDurationHookFunc())))
		if err != nil {
			return fmt.Errorf("unmarshal viper config: %v", err)
		}

		if config.DatabasePath == "" {
			// DB path could not be found. Re-run defaultDataDir to get the underlying error and return it.
			dbPath, err := defaultDataDir()
			if err != nil {
				return fmt.Errorf("find data directory: %v", err)
			}
			// Somehow, defaultDataDir() returned an error the first time but succeeded this time.
			// This is very unlikely, but we handle it anyway.
			config.DatabasePath = dbPath
		}
		return nil
	}

	defaultDBPath, err := defaultDataDir()
	if err != nil {
		// We can't return an error here because `--help` would fail.
		// We check that the db path is not empty in PreRunE.
		defaultDBPath = ""
	}

	// For testing purposes, these variables cannot be declared outside the function because Cobra
	// may mutate their values.
	defaultLogLevel := utils.INFO
	defaultNetwork := utils.MAINNET

	junoCmd.Flags().StringVar(&cfgFile, configF, defaultConfig, configFlagUsage)
	junoCmd.Flags().Var(&defaultLogLevel, logLevelF, logLevelFlagUsage)
	junoCmd.Flags().Bool(httpF, defaultHTTP, httpUsage)
	junoCmd.Flags().String(httpHostF, defaulHost, httpHostUsage)
	junoCmd.Flags().Uint16(httpPortF, defaultHTTPPort, httpPortUsage)
	junoCmd.Flags().Bool(wsF, defaultWS, wsUsage)
	junoCmd.Flags().String(wsHostF, defaulHost, wsHostUsage)
	junoCmd.Flags().Uint16(wsPortF, defaultWSPort, wsPortUsage)
	junoCmd.Flags().String(dbPathF, defaultDBPath, dbPathUsage)
	junoCmd.Flags().Var(&defaultNetwork, networkF, networkUsage)
	junoCmd.Flags().String(ethNodeF, defaultEthNode, ethNodeUsage)
	junoCmd.Flags().Bool(pprofF, defaultPprof, pprofUsage)
	junoCmd.Flags().String(pprofHostF, defaulHost, pprofHostUsage)
	junoCmd.Flags().Uint16(pprofPortF, defaultPprofPort, pprofPortUsage)
	junoCmd.Flags().Bool(colourF, defaultColour, colourUsage)
	junoCmd.Flags().Duration(pendingPollIntervalF, defaultPendingPollInterval, pendingPollIntervalUsage)
	junoCmd.Flags().Bool(p2pF, defaultP2p, p2pUsage)
	junoCmd.Flags().String(p2pAddrF, defaultP2pAddr, p2PAddrUsage)
	junoCmd.Flags().String(p2pBootPeersF, defaultP2pBootPeers, p2pBootPeersUsage)
	junoCmd.Flags().Bool(metricsF, defaultMetrics, metricsUsage)
	junoCmd.Flags().String(metricsHostF, defaulHost, metricsHostUsage)
	junoCmd.Flags().Uint16(metricsPortF, defaultMetricsPort, metricsPortUsage)
	junoCmd.Flags().Bool(grpcF, defaultGRPC, grpcUsage)
	junoCmd.Flags().String(grpcHostF, defaulHost, grpcHostUsage)
	junoCmd.Flags().Uint16(grpcPortF, defaultGRPCPort, grpcPortUsage)

	return junoCmd
}

func defaultDataDir() (string, error) {
	const junoDirName = "juno"

	if userDataDir := os.Getenv("XDG_DATA_HOME"); userDataDir != "" {
		return filepath.Join(userDataDir, junoDirName), nil
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	} else if userHomeDir == "" {
		return "", errors.New("home directory not found")
	}
	return filepath.Join(userHomeDir, ".local", "share", junoDirName), nil
}
