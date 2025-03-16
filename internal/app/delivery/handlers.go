package delivery

import (
	"awesomeProject5/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	uc IUseCase
}

func New(ucase IUseCase) *Handler {
	return &Handler{uc: ucase}
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task entity.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.uc.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetTasks(c *gin.Context) {
	statusStr := c.Query("status")
	priorityStr := c.Query("priority")

	c.Header("Cache-Control", "public, max-age=3600")

	if statusStr == "" && priorityStr == "" {
		//c.JSON(http.StatusOK, repository.Storage)
		return
	}

	var tasks []entity.Task

	//for _, task := range repository.Storage {
	//	if statusStr != "" {
	//		status, err := strconv.ParseBool(statusStr)
	//		if err != nil || task.Status != status {
	//			continue
	//		}
	//	}
	//	if priorityStr != "" {
	//		priority, err := strconv.Atoi(priorityStr)
	//		if err != nil || task.Priority != priority {
	//			continue
	//		}
	//	}
	//
	//	tasks = append(tasks, task)
	//}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	//parsedID, err := uuid.Parse(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//task, exists := repository.Storage[parsedID]
	//if !exists {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	//	return
	//}
	//
	//if err = c.ShouldBindJSON(&task); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//repository.Storage[parsedID] = task
}

func (h *Handler) DeleteTask(c *gin.Context) {
	//parsedID, err := uuid.Parse(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//if _, exists := repository.Storage[parsedID]; !exists {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	//	return
	//}
	//
	//delete(repository.Storage, parsedID)
}
