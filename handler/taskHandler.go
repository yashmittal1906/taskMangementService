package handler

import (
    "net/http"
    "taskMangementService/model"
    "taskMangementService/service"
    "github.com/gin-gonic/gin"
)

type TaskHandler struct {
    Service *service.TaskService
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
    tasks := r.Group("/tasks")
    {
        tasks.POST("/", h.CreateTask)
    }
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
    var task model.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Service.CreateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, task)
}
