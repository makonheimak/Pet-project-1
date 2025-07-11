package main

import (
	"log"
	"myproject/internal/db"
	"myproject/internal/handlers"
	"myproject/internal/taskservice"

	"myproject/internal/web/tasks"

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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
