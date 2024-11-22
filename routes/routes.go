package routes

import (
	"todo-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    // Route group for To-Do
    todoRoutes := r.Group("/todos")
    {
        todoRoutes.POST("", handlers.CreateTodoHandler)  // Create
        todoRoutes.GET("", handlers.GetTodosHandler)     // Read all
        todoRoutes.PUT("/:id", handlers.UpdateTodoHandler) // Update
        todoRoutes.DELETE("/:id", handlers.DeleteTodoHandler) // Delete
    }

    return r
}
