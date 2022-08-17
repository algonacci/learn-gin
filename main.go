package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Index struct {
	Message string `json:"message"`
}

var index = Index{Message: "Hello World"}

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, index)
}

type Person struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}

var persons = []Person{
	{ID: "1", FullName: "John Doe", Age: 25},
	{ID: "2", FullName: "Jane Doe", Age: 27},
}

func getPerson(c *gin.Context) {
	c.JSON(http.StatusOK, persons)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")
	for _, person := range persons {
		if person.ID == id {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "Person not found"})
}

func createPerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	persons = append(persons, person)
	c.JSON(http.StatusOK, person)
	if err := c.Errors.ByType(gin.ErrorTypeBind).String(); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", GetIndex)
	router.GET("/persons", getPerson)
	router.GET("/persons/:id", getPersonByID)
	router.POST("/persons", createPerson)
	router.Run()
}
