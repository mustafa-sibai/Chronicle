package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"github.com/mustafa-sibai/chronicle/registration-service/handlers"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_RegistrationService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	mongodb.Connect("mongodb://localhost:27017")
	valkeydb.Connect("localhost:6379")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/register", handlers.RegisterHandler)

	fmt.Println("Starting service on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
