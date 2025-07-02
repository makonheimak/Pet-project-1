package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", postTask)
	e.GET("/task", getTasks)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	e.Start("localhost:8080")
}

var tasks = []Task{}

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}

	tasks = append(tasks, req)

	return c.JSON(http.StatusOK, map[string]string{"status": "task saved"})
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Task = req.Task
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "task not found"})
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}
