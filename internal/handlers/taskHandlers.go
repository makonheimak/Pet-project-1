package handlers

import (
	taskservice "myproject/internal/taskService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandlers struct {
	service taskservice.TaskService
}

func NewTaskHandlers(s taskservice.TaskService) *TaskHandlers {
	return &TaskHandlers{service: s}
}

func (h *TaskHandlers) PostTasks(c echo.Context) error {
	var req taskservice.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}

	task, err := h.service.CreateTask(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not Create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandlers) GetTaskss(c echo.Context) error {
	tasks, err := h.service.GetAllTask()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandlers) PatchTasks(c echo.Context) error {
	id := c.Param("id")

	var req taskservice.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	updateTask, err := h.service.UpdateTask(id, req.Task)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, updateTask)
}

func (h *TaskHandlers) DeleteTasks(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get task"})
	}

	return c.NoContent(http.StatusNoContent)
}
