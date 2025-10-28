package config

import (
	"github.com/go-ini/ini"
)

// Config represents configuration variables
type Config struct {
	ApiKey    string
	ApiSecret string
	Subdomain string
	IsProd    bool
	DevUrl    string
}

// LoadConfig gets config from provided path
func LoadConfig(configPath string) (lsConfig Config, err error) {
	err = ini.MapTo(&lsConfig, configPath)
	return
}
