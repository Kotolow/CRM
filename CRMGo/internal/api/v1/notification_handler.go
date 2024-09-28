package v1

//type NotificationHandler struct {
//	taskService database.NotificationService
//}
//
//func NewNotificationHandler(service database.NotificationService) *NotificationHandler {
//	return &GanttHandler{taskService: service}
//}
//
//// GetChart godoc
//// @Summary      Get Gantt Chart
//// @Description  Get All Info for Gantt Chart
//// @Tags         charts
//// @Produce      json
//// @Param        code   path   string true  "Project code"
//// @Success      200  {object}  response.Response
//// @Failure      500  {object}  response.Response
//// @Router       /v1/charts/{code} [get]
//func (h *GanttHandler) GetChart(c *gin.Context) {
//	taskResponse, err := h.taskService.GetGanttDataByProject(c.Param("code"))
//	if err != nil {
//		utils.InternalError(c, err)
//		return
//	}
//
//	webResponse := response.Response{
//		Code:   200,
//		Status: http.StatusText(http.StatusOK),
//		Data:   taskResponse,
//	}
//
//	c.Header("Content-Type", "application/json")
//	c.JSON(http.StatusOK, webResponse)
//}
