package main

import (
	"context"
	"fmt"
	"github.com/mustafa-sibai/chronicle/authentication-service/handlers"
	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_AuthenticationService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	mongodb.Connect("mongodb://localhost:27017/?replicaSet=rs0")
	valkeydb.Connect("localhost:6379")

	if err := collections.IndexRefreshToken(context.Background()); err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Post("/token/refresh", handlers.RefreshHandler)
	r.Post("/logout", handlers.LogoutHandler)

	fmt.Println("Starting service on :3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
