package livespace_client

import (
	"github.com/arturwwl/livespace-golang-client/config"
	//p24Error "github.com/arturwwl/p24-golang-client/error"
)

type LivespaceClient struct {
	Config config.Config
}

func New(configPath string) (client *LivespaceClient, err error) {
	client = new(LivespaceClient)
	client.Config, err = config.LoadConfig(configPath)

	return
}
