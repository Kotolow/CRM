package utils

import (
	"CRMGo/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
)

func InternalError(c *gin.Context, err error) {
	slog.Error(err.Error())
	webResponse := response.Response{
		Code:   500,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   nil,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
	return
}

func NewTaskName(taskName string) (string, error) {
	re := regexp.MustCompile(`(.*?)(\d+)$`)

	matches := re.FindStringSubmatch(taskName)

	number, err := strconv.Atoi(matches[2])
	if err != nil {
		return "", err
	}
	fmt.Println("number: ", number)
	newTaskName := fmt.Sprintf("%s%d", matches[1], number+1)

	return newTaskName, nil
}
