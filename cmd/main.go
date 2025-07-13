package main

import (
	"log"
	"myproject/internal/db"
	"myproject/internal/handlers"
	"myproject/internal/taskservice"
	"myproject/internal/userservice"

	"myproject/internal/web/tasks"
	"myproject/internal/web/users"

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

	userRepo := userservice.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
