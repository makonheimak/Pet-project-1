package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var task string

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", postTask)
	e.GET("/task", getTask)

	e.Start("localhost:8080")

}

type Task struct {
	Task string `json:"task"`
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}

	task = req.Task
	return c.JSON(http.StatusOK, map[string]string{"status": "task saved"})

}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "hell привет, " + task})
}
