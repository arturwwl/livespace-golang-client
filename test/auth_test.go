package test

import (
	livespace_client "github.com/arturwwl/livespace-golang-client/client"
	"testing"
)

func TestAuth(t *testing.T) {
	client, err := livespace_client.New("conf/cfg.ini")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetAuth()
	if err != nil {
		t.Fatal(err)
	}
}
