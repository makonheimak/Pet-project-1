package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", postTask)
	e.GET("/task", getTasks)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	e.Start("localhost:8080")
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type Task struct {
	ID   string `gorm: "primaryKey" json:"id"`
	Task string `json:"task"`
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}

	if err := db.Create(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add tasks"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "task saved"})
}

func getTasks(c echo.Context) error {
	var tasks = []Task{}

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	if req.Task != "" {
		task.Task = req.Task
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
