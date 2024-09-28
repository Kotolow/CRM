package gantt

import "time"

type GanttChartData struct {
	ProjectCode string      `json:"project_code"`
	ProjectName string      `json:"project_name"`
	Tasks       []GanttTask `json:"tasks"`
}

type GanttTask struct {
	TaskID     string    `json:"task_id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Priority   string    `json:"priority"`
	AssignedTo string    `json:"assigned_to"`
	StartDate  time.Time `json:"start_date"`
	DueDate    time.Time `json:"due_date"`
	TimeSpent  int       `json:"time_spent"`
}
