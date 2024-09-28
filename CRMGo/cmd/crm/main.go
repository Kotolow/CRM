package main

import (
	"CRMGo/internal/api"
	v1 "CRMGo/internal/api/v1"
	"CRMGo/pkg/database"
	"CRMGo/pkg/git"
	"os"
)

func main() {

	token := os.Getenv("GIT_TOKEN")

	db := database.DatabaseConnection()

	taskRepo := database.NewTaskRepo(db)
	projectRepo := database.NewProjectRepo(db)
	ganttRepo := database.NewGanttChartRepo(db)
	taskHandler := v1.NewTaskHandler(*taskRepo)
	projectHandler := v1.NewProjectHandler(*projectRepo)
	ganttHandler := v1.NewGanttHandler(*ganttRepo)
	gitService := git.NewGitHubRepo(token)
	gitHandler := v1.NewGitHubHandler(*gitService)
	controller := v1.NewCRMController(*taskHandler, *projectHandler, *gitHandler, *ganttHandler)

	router := api.NewRouter(controller)

	err := router.Run(":8081")
	//err := router.Run()
	if err != nil {
		panic(err)
	}

}
