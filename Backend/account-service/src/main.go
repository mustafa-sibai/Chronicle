package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/account-service/handlers"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_AccountService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	mongodb.Connect("mongodb://localhost:27017/?replicaSet=rs0")
	valkeydb.Connect("localhost:6379")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/update/email", handlers.UpdateEmailHandler)
	r.Post("/update/password", handlers.UpdatePasswordHandler)
	r.Post("/update/username", handlers.UpdateUsernameHandler)

	fmt.Println("Starting service on :3003")
	log.Fatal(http.ListenAndServe(":3003", r))
}
