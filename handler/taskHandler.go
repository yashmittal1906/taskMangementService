package handler

import (
	"net/http"
	"strconv"
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
        tasks.GET("/", h.GetTasks)
        tasks.GET("/:id", h.GetTask)
        tasks.PUT("/:id", h.UpdateTask)
        tasks.DELETE("/:id", h.DeleteTask)
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

func (h *TaskHandler) GetTasks(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    status := c.Query("status")

    tasks, err := h.Service.GetTasks(status, page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    task, err := h.Service.GetTaskByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var task model.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = uint(id)
    if err := h.Service.UpdateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Service.DeleteTask(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Deleted the given task"})
}
