package settings

import (
	"os"
)

// Settings Represents the settings for the manager
type Settings struct {
	HostName string `json:"host_name"`
	Port     string `json:"port"`
}

// FromEnvironment returns a Settings object from an environment field
func FromEnvironment() (Settings, error) {
	settings := Settings{}
	name, err := os.Hostname()
	if err != nil {
		return settings, err
	}
	settings.HostName = name
	settings.Port = "8081"

	return settings, nil
}
