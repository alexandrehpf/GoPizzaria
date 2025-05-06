package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Toscana", Preco: 49.5},
	{ID: 2, Nome: "Marguerita", Preco: 79.5},
	{ID: 3, Nome: "Atum com Queijo", Preco: 69.5},
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.POST("/pizzas", postPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, gin.H{"pizza": pizza})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}
