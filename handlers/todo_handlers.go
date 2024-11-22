package handlers

import (
	"net/http"
	"todo-service/config"
	"todo-service/models"

	"github.com/gin-gonic/gin"
)

// Create a new To-Do
func CreateTodoHandler(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }

    c.JSON(http.StatusCreated, todo)
}

// Get all To-Do items
func GetTodosHandler(c *gin.Context) {
    var todos []models.Todo
    if err := config.DB.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
        return
    }

    c.JSON(http.StatusOK, todos)
}

// Update a To-Do
func UpdateTodoHandler(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")

    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

// Delete a To-Do
func DeleteTodoHandler(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")

    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    if err := config.DB.Delete(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
