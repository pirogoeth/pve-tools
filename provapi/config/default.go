package config

func DefaultConfig() *Config {
	return &Config{
		ListenAddress: ":8080",
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
