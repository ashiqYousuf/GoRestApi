package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Title: "Golang Learning", Completed: true},
	{ID: "2", Title: "Have a dinner", Completed: true},
	{ID: "3", Title: "Going Boh", Completed: false},
	{ID: "4", Title: "Reading Arabic", Completed: false},
}

func getTodos(c *gin.Context) {
	// ?c contains information about the incoming Http request
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": "false"})
		return
	}
	c.IndentedJSON(http.StatusOK, *todo)
}

func toggleStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"success": "false"})
		return
	}
	(*todo).Completed = !(*todo).Completed
	c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	updatedTodo := []Todo{}
	for _, todo := range todos {
		if todo.ID == id {
			continue
		}
		updatedTodo = append(updatedTodo, todo)
	}
	todos = updatedTodo
	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}

func main() {
	// ?Create a server
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleStatus)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}
