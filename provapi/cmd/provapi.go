package main

import (
	"context"
	"net"

	"github.com/gin-gonic/gin"
	"tailscale.com/tsnet"

	"github.com/pirogoeth/pve-tools/pkg/config"
	"github.com/pirogoeth/pve-tools/pkg/logging"
	"github.com/pirogoeth/pve-tools/pkg/utils"
	"github.com/pirogoeth/pve-tools/provapi/api"
	"github.com/pirogoeth/pve-tools/provapi/clients"
	apiCfg "github.com/pirogoeth/pve-tools/provapi/config"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	logging.New()

	// Load config
	cfg, err := config.LoadWithUnderlay(apiCfg.DefaultConfig())
	if err != nil {
		logging.Fatalf("Failed to load config: %s", err)
	}

	// Create a Gin engine
	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(logging.LoggerWriter()))
	engine.Use(gin.Recovery())

	if err := clients.Init(ctx, cfg); err != nil {
		logging.Fatalf("Failed to initialize clients: %s", err)
	}
	if err := api.Init(ctx, cfg, engine); err != nil {
		logging.Fatalf("Failed to initialize API: %s", err)
	}

	// Start serving
	if cfg.Tailscale.Enabled {
		// Start a Tailscale listener
		listener, err := primeTailscaleServer(ctx, cfg)
		if err != nil {
			logging.Fatalf("Failed to start Tailscale listener: %s", err)
		}
		engine.RunListener(listener)
	} else {
		engine.Run(cfg.ListenAddress)
	}
}

func primeTailscaleServer(ctx context.Context, cfg *apiCfg.Config) (net.Listener, error) {
	// Create a custom, trace-level logger for Tailscale.
	logf := logging.GetLogger().WithField("component", "tailscale").Tracef
	s := &tsnet.Server{
		Hostname: cfg.Tailscale.Hostname,
		Logf:     logf,
	}
	go utils.OnDone(ctx, func() { s.Close() })
	if err := s.Start(); err != nil {
		return nil, err
	}

	listener, err := s.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return nil, err
	}

	return listener, nil
}
