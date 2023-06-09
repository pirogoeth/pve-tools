package config

import "encoding/json"

type Config struct {
	// Database is the configuration for the database.
	Database DatabaseCfg `json:"database"`
	// ListenAddress specifies the address to listen on.
	ListenAddress string `json:"listen_address"`
	// Tailscale is the configuration for Tailscale.
	Tailscale TailscaleCfg `json:"tailscale"`
	// Proxmox is the configuration for the Proxmox client.
	Proxmox ProxmoxCfg `json:"proxmox"`
	// ImageFamilies is the list of image families configs.
	ImageFamilies ImageFamilyCfgs `json:"image_families"`
}

type DatabaseCfg struct {
	// Type is the type of database to use. Currently only "sqlite3" is supported.
	Type string `json:"type"`
	// Config is the configuration for the connected database.
	Config *json.RawMessage `json:"config"`
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
	UserId string `json:"api_user_id"`
	// ApiToken is the token to use for authenticating
	ApiToken string `json:"api_token"`
	// TaskTimeoutSeconds is the number of seconds to wait for a task to complete
	TaskTimeoutSeconds int `json:"task_timeout_seconds"`
	// TlsInsecure specifies whether to skip TLS verification
	TlsInsecure bool `json:"tls_insecure"`
}

type ImageFamilyCfg struct {
	// Name is the name of the image family.
	Name string `json:"name"`
	// Description is the description of the image family.'
	Description string `json:"description"`
	// ImagePattern is the pattern used to match images to this family.
	ImagePattern string `json:"image_pattern"`
}

type ImageFamilyCfgs []ImageFamilyCfg
