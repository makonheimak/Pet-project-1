package main

import (
	"log"
	"myproject/internal/db"
	"myproject/internal/handlers"
	taskservice "myproject/internal/taskService"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandlers(taskService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", taskHandlers.PostTask)
	e.GET("/task", taskHandlers.GetTasks)
	e.PATCH("/task/:id", taskHandlers.PatchTask)
	e.DELETE("/task/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
