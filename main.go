package main

import (
	"fmt"
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Nome: "Toscana", Preco: 49.5},
		{ID: 2, Nome: "Marguerita", Preco: 79.5},
		{ID: 3, Nome: "Atum com Queijo", Preco: 69.5},
	}
	fmt.Println(pizzas)
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
