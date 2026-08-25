package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/authentication-server/handlers"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationNameAuthenticationServer,
		EnvironmentType: common.EnvironmentTypeDevelopment,
	})

	mongodb.Connect("mongodb://localhost:27017")
	valkeydb.Connect("localhost:6379")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/authenticate", handlers.AuthenticateHandler)

	fmt.Println("Starting server on :3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
