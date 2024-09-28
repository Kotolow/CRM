package main

import (
	"CRMGo/internal/api"
	v1 "CRMGo/internal/api/v1"
	"CRMGo/pkg/database"
	"CRMGo/pkg/git"
	"os"
)

func main() {
	//os.Setenv("DB_HOST", "0.tcp.eu.ngrok.io")
	//os.Setenv("DB_PORT", "10717")
	//os.Setenv("DB_USER", "root")
	//os.Setenv("DB_PASSWORD", "password")
	//os.Setenv("DB_NAME", "root")
	//os.Setenv("SMTP_HOST", "172.19.0.2")
	//os.Setenv("SMTP_PORT", "1025")
	//os.Setenv("GIT_NAME", "HAndHTemp")
	//os.Setenv("GIT_TOKEN", "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")

	token := os.Getenv("GIT_TOKEN")

	db := database.DatabaseConnection()

	taskRepo := database.NewTaskRepo(db)
	projectRepo := database.NewProjectRepo(db)
	ganttRepo := database.NewGanttChartRepo(db)
	//notificationRepo := database.NewNotificationService(db)
	taskHandler := v1.NewTaskHandler(*taskRepo)
	projectHandler := v1.NewProjectHandler(*projectRepo)
	ganttHandler := v1.NewGanttHandler(*ganttRepo)
	gitService := git.NewGitHubRepo(token)
	gitHandler := v1.NewGitHubHandler(*gitService)
	controller := v1.NewCRMController(*taskHandler, *projectHandler, *gitHandler, *ganttHandler)

	router := api.NewRouter(controller)

	//err := router.Run(":8081")
	err := router.Run()
	if err != nil {
		panic(err)
	}

}
