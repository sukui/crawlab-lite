package controllers

import (
	"crawlab-lite/constants"
	"crawlab-lite/forms"
	"crawlab-lite/services"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetTaskList(c *gin.Context) {
	var page forms.TaskPageForm

	if err := c.ShouldBindQuery(&page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if total, tasks, err := services.QueryTaskPage(page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccessList(c, total, tasks)
	}
}

func GetTask(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		HandleError(http.StatusBadRequest, c, errors.New("invalid task id"))
		return
	}

	if task, err := services.QueryTaskById(id); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		if task == nil {
			HandleError(http.StatusNotFound, c, errors.New("task not found"))
			return
		}
		HandleSuccess(c, task)
	}
}

func CreateTask(c *gin.Context) {
	var form forms.TaskForm

	if err := c.ShouldBindJSON(&form); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if res, err := services.AddTask(form); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}

func DeleteTask(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		HandleError(http.StatusBadRequest, c, errors.New("invalid task id"))
		return
	}

	if res, err := services.RemoveTask(id); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}

func UpdateTaskCancel(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		HandleError(http.StatusBadRequest, c, errors.New("invalid task id"))
		return
	}

	if res, err := services.CancelTask(id, constants.TaskStatusCancelled); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}

func PostTaskRestart(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		HandleError(http.StatusBadRequest, c, errors.New("invalid task id"))
		return
	}

	if res, err := services.RestartTask(id); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}

func GetTaskLogList(c *gin.Context) {
	var page forms.TaskLogPageForm

	if err := c.ShouldBindUri(&page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if err := c.ShouldBindQuery(&page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if _, err := uuid.FromString(page.TaskId); err != nil {
		HandleError(http.StatusBadRequest, c, errors.New("invalid task id"))
		return
	}

	if total, results, err := services.QueryTaskLogPage(page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccessList(c, total, results)
	}
}
