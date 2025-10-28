package livespaceclient

import (
	"github.com/arturwwl/livespace-golang-client/config"
)

type LivespaceClient struct {
	Config config.Config
}

func New(configPath string) (client *LivespaceClient, err error) {
	client = new(LivespaceClient)
	client.Config, err = config.LoadConfig(configPath)

	return
}
