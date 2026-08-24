package valkeydb

import (
	"context"
	"fmt"
	"github.com/valkey-io/valkey-go"
)

func Connect(valkeyURI string) {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{valkeyURI}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		panic(err)
	}

	// client.Do(ctx, client.B().Set().Key("xx").Value("ll").Nx().Build()).Error()

	fmt.Println("Connected to Valkey")
}
