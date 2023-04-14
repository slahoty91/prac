package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todoList = []todo{
	{ID: "1", Item: "Tets", Completed: true},
	{ID: "2", Item: "Tets1", Completed: false},
	{ID: "3", Item: "Tets2", Completed: false},
}

func getToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todoList)
}

func addToDo(context *gin.Context) {
	var newToDo todo
	if err := context.BindJSON(&newToDo); err != nil {
		return
	}
	todoList = append(todoList, newToDo)
	context.IndentedJSON(http.StatusCreated, newToDo)
}

func getToDobyID(id string) (*todo, error) {
	for i, t := range todoList {
		if t.ID == id {
			return &todoList[i], nil
		}
	}
	return nil, errors.New("not found")
}

func getToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getToDobyID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getToDobyID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"msg": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.GET("/todos/:id", getToDo)
	router.POST("/addTodos", addToDo)
	router.PATCH("/todo/:id", toggleToDo)
	router.Run("localhost:9090")
}
