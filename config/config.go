package config

import (
	"github.com/go-ini/ini"
)

type Config struct {
	ApiKey    string
	ApiSecret string
	Subdomain string
	IsProd    bool
	DevUrl    string
}

func LoadConfig(configPath string) (lsConfig Config, err error) {
	err = ini.MapTo(&lsConfig, configPath)
	return
}
