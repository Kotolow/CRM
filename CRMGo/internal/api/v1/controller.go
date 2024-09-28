package v1

type CRMController struct {
	TaskHandler    TaskHandler
	ProjectHandler ProjectHandler
	GitHubHandler  GitHubHandler
	GanttHandler   GanttHandler
}

func NewCRMController(taskHandler TaskHandler, projectHandler ProjectHandler, gitHubHandler GitHubHandler, ganttHandler GanttHandler) *CRMController {
	return &CRMController{
		TaskHandler:    taskHandler,
		ProjectHandler: projectHandler,
		GitHubHandler:  gitHubHandler,
		GanttHandler:   ganttHandler,
	}
}
