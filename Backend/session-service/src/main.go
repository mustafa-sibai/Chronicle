package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"github.com/mustafa-sibai/chronicle/session-service/handlers"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_SessionService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	valkeydb.Connect("localhost:6379")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/create", handlers.CreateHandler)
	r.Post("/destroy", handlers.DestroyHandler)

	fmt.Println("Starting service on :3002")
	log.Fatal(http.ListenAndServe(":3002", r))
}
