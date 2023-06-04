package clients

import (
	"context"
	"fmt"

	"github.com/Telmate/proxmox-api-go/proxmox"

	"github.com/pirogoeth/pve-tools/pkg/logging"
	apiCfg "github.com/pirogoeth/pve-tools/provapi/config"
)

var (
	clients struct {
		proxmox *proxmox.Client
	}
)

func Proxmox() *proxmox.Client {
	return clients.proxmox
}

func Init(ctx context.Context, cfg *apiCfg.Config) error {
	clients := new(struct {
		proxmox *proxmox.Client
	})

	var err error
	clients.proxmox, err = initProxmox(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize Proxmox client: %w", err)
	}

	return nil
}

func initProxmox(cfg *apiCfg.Config) (*proxmox.Client, error) {
	pmCfg := cfg.Proxmox
	pmClient, err := proxmox.NewClient(pmCfg.ApiUrl, nil, "", nil, "", pmCfg.TaskTimeoutSeconds)
	if err != nil {
		return nil, err
	}

	pmClient.SetAPIToken(pmCfg.UserId, pmCfg.ApiToken)
	versionResp, err := pmClient.GetVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to check Proxmox API version during client init: %w", err)
	}

	logging.Infof("Proxmox client initialized - API version %#v", versionResp)

	return pmClient, nil
}
