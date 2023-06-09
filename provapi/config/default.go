package config

import "encoding/json"

func DefaultConfig() *Config {
	rawDbCfg := json.RawMessage(`{"path": "provapi.db"}`)

	return &Config{
		ListenAddress: ":8080",
		Database: DatabaseCfg{
			Type:   "sqlite3",
			Config: &rawDbCfg,
		},
		Tailscale: TailscaleCfg{
			Enabled:  false,
			Hostname: "pveprovapi",
		},
		Proxmox: ProxmoxCfg{
			ApiUrl:             "",
			UserId:             "",
			ApiToken:           "",
			TaskTimeoutSeconds: 30,
		},
	}
}
