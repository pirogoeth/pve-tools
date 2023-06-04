package config

import (
	"bytes"
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

const (
	ENV_CONFIG_SOURCE = "CONFIG_SOURCE"
	ENV_CONFIG_TYPE   = "CONFIG_TYPE"
	CONFIG_TYPE_FILE  = "file"
)

// Load loads the configuration from the source specified by ENV_CONFIG_SOURCE.
// Expects any backend-specific configuration to be provided in the environment.
// If ENV_CONFIG_TYPE is empty, it defaults to "file".
func Load[T any]() (*T, error) {
	return LoadWithUnderlay[T](nil)
}

// LoadWithUnderlay loads the configuration from the source specified by ENV_CONFIG_SOURCE.
// Expects any backend-specific configuration to be provided in the environment.
// If ENV_CONFIG_TYPE is empty, it defaults to "file".
// The "underlay" effectively serves as a default configuration with the true source loading over it.
// Note that this only works if `omitempty` is set on the struct fields.
func LoadWithUnderlay[T any](underlay *T) (*T, error) {
	configType := os.Getenv(ENV_CONFIG_TYPE)
	if configType == "" {
		configType = CONFIG_TYPE_FILE
	}

	switch configType {
	case CONFIG_TYPE_FILE:
		return loadConfigFromFile[T](underlay)
	default:
		return nil, fmt.Errorf("Unknown config type: %s", configType)
	}
}

func loadConfigFromFile[T any](underlay *T) (*T, error) {
	if underlay == nil {
		underlay = new(T)
	}

	configSource := os.Getenv(ENV_CONFIG_SOURCE)
	if configSource == "" {
		return nil, fmt.Errorf("CONFIG_SOURCE is not set")
	}

	file, err := os.Open(configSource)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bytes.NewBufferString("")
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf.Bytes(), &underlay)
	if err != nil {
		return nil, err
	}

	return underlay, nil
}
