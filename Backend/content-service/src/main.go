package main

import (
	"fmt"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/router"
	"github.com/mustafa-sibai/chronicle/content-service/handlers"
	"log"
	"net/http"
)

func main() {
	common.SetConfig(common.Config{
		ApplicationName: common.ApplicationName_ContentService,
		EnvironmentType: common.EnvironmentType_Development,
	})

	mongodb.Connect("mongodb://localhost:27017/?replicaSet=rs0")

	r := router.NewRouter()

	r.Get("/health", handlers.HealthHandler)
	r.Post("/itemTemplates/create", handlers.CreateItemTemplateHandler)
	r.Post("/itemTemplates/get", handlers.GetItemTemplateHandler)
	r.Post("/itemTemplates/list", handlers.ListItemTemplatesHandler)
	r.Post("/itemTemplates/delete", handlers.DeleteItemTemplateHandler)
	r.Post("/quests/create", handlers.CreateQuestHandler)
	r.Post("/quests/get", handlers.GetQuestHandler)
	r.Post("/quests/list", handlers.ListQuestsHandler)
	r.Post("/quests/delete", handlers.DeleteQuestHandler)

	fmt.Println("Starting service on :3005")
	log.Fatal(http.ListenAndServe(":3005", r))
}
