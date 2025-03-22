package delivery

import (
	"awesomeProject5/internal/entity"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	uc IUseCase
}

func New(ucase IUseCase) *Handler {
	return &Handler{uc: ucase}
}

func (h *Handler) CreateTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task entity.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.uc.CreateTask(ctx, task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetTasks(c *gin.Context) {
}

func (h *Handler) UpdateTask(c *gin.Context) {
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
