package app

import (
	"awesomeProject5/internal/app/delivery"
	"github.com/gin-gonic/gin"
)

func InitEndPoints(engine *gin.Engine, handler *delivery.Handler) {
	engine.POST("/task", handler.CreateTask)       // POST http://localhost:8080/task
	engine.GET("/task", handler.GetTasks)          // GET http://localhost:8080/task
	engine.PUT("/task/:id", handler.UpdateTask)    // PUT http://localhost:8080/task
	engine.DELETE("/task/:id", handler.DeleteTask) // DELETE http://localhost:8080/task
}
