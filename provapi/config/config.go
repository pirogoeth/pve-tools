package config

type Config struct {
	// ListenAddress specifies the address to listen on.
	ListenAddress string       `json:"listen_address"`
	Tailscale     TailscaleCfg `json:"tailscale"`
	Proxmox       ProxmoxCfg   `json:"proxmox"`
}

type TailscaleCfg struct {
	// Enabled specifies whether the service should be bound to a Tailscale Tailnet.
	Enabled bool `json:"enabled,omitempty"`
	// Hostname specifies the name this service should advertise on the Tailnet.
	Hostname string `json:"hostname,omitempty"`
}

type ProxmoxCfg struct {
	// ApiUrl is the URL of the Proxmox API
	ApiUrl string `json:"api_url"`
	// UserId is the ID of the token to use for authenticating.
	// In the form of "<USER>@<REALM>!<TOKENID>"
	UserId string `json:"user_id"`
	// ApiToken is the token to use for authenticating
	ApiToken string `json:"token"`
	// TaskTimeoutSeconds is the number of seconds to wait for a task to complete
	TaskTimeoutSeconds int `json:"task_timeout_seconds"`
}
