package main

import (
	"taskMangementService/handler"
	"taskMangementService/model"
	"taskMangementService/repository"
	"taskMangementService/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to the database")
    }

    db.AutoMigrate(&model.Task{})

    repo := &repository.TaskRepository{DB: db}
    taskService := &service.TaskService{Repo: repo}
    taskHandler := &handler.TaskHandler{Service: taskService}

    r := gin.Default()
    taskHandler.RegisterRoutes(r)

    r.Run(":8081")
}
