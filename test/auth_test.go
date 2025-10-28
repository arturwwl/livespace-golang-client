package test

import (
	livespaceclient "github.com/arturwwl/livespace-golang-client/client"
	"testing"
)

func TestAuth(t *testing.T) {
	client, err := livespaceclient.New("conf/cfg.ini")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetAuth()
	if err != nil {
		t.Fatal(err)
	}
}
