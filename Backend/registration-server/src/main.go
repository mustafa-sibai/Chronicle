package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/registration-server/handlers"
	"github.com/mustafa-sibai/chronicle/registration-server/router"
	"log"
	"net/http"
)

func main() {
	mongodb.Connect("mongodb://localhost:27017")
	valkeydb.Connect("localhost:6379")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)

	fmt.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
