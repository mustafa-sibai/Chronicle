package valkeydb

import (
	"context"
	"fmt"
	"github.com/valkey-io/valkey-go"
)

var client valkey.Client

func Connect(valkeyURI string) {
	c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{valkeyURI}})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	if err := c.Do(ctx, c.B().Ping().Build()).Error(); err != nil {
		panic(err)
	}

	client = c

	fmt.Println("Connected to Valkey")
}

func GetValkeyClient() valkey.Client {
	return client
}
