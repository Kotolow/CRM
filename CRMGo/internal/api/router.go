package api

import (
	_ "CRMGo/docs"
	"CRMGo/internal/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Horns and Hooves CRM API
//	@version		1.0
//	@description	This is a CRM doc
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func NewRouter(controller *v1.CRMController) *gin.Engine {
	service := gin.Default()

	//service.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*", "http://localhost:3001", "http://127.0.0.1:3001", "http://localhost:3000", "http://127.0.0.1:3000"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	service.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1Router := service.Group("/v1")

	chartsRouter := v1Router.Group("/charts")
	chartsRouter.GET("/:code", controller.GanttHandler.GetChart)

	gitRouter := v1Router.Group("/git")
	gitRouter.POST("/:owner/:repo/new_branch", controller.GitHubHandler.CreateBranch)
	gitRouter.POST("/:owner/:repo/new_pr", controller.GitHubHandler.CreatePullRequest)
	gitRouter.GET("/repos", controller.GitHubHandler.GetRepos)
	gitRouter.GET("/:taskId/branches", controller.GitHubHandler.GetBranches)
	gitRouter.GET("/:taskId/commits", controller.GitHubHandler.GetCommits)
	gitRouter.GET("/:taskId/pull_requests", controller.GitHubHandler.GetPullRequests)

	projectRouter := v1Router.Group("/projects")
	projectRouter.GET("", controller.ProjectHandler.FindAll)
	projectRouter.GET("/:code", controller.ProjectHandler.FindById)
	projectRouter.POST("", controller.ProjectHandler.Create)
	projectRouter.PUT("/:code", controller.ProjectHandler.Update)
	projectRouter.DELETE("/:code", controller.ProjectHandler.Delete)

	taskRouter := projectRouter.Group("/:code/tasks")
	taskRouter.GET("", controller.TaskHandler.FindAll)
	taskRouter.GET("/:taskId", controller.TaskHandler.FindById)
	taskRouter.POST("", controller.TaskHandler.Create)
	taskRouter.PUT("/:taskId", controller.TaskHandler.Update)
	taskRouter.DELETE("/:taskId", controller.TaskHandler.Delete)

	return service
}
