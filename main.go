package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// type user struct {
// 	Name string `json:"name"`

// }

type pet struct {
	Name string  `json:"name" binding:"required"`
	Food float64 `json:"food" binding:"required"`
}

var pets = []pet{
	{Name: "silly", Food: 100.00},
}

func main() {
	router := gin.Default()
	router.GET("/pets", getPets)
	router.POST("/create", createPet)
	go updatePets(pets)
	router.Run()
}

func getPets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pets)
}

func createPet(c *gin.Context) {
	var newPet pet

	if err := c.Bind(&newPet); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pets = append(pets, newPet)
	c.JSON(http.StatusCreated, newPet)
}

func updatePets(ps []pet) []pet {
	for range time.Tick(time.Second * 5) {
		for m, p := range ps {
			p.Food -= 1
			ps[m] = p
			fmt.Println("works?", p.Food)
		}
	}
	return pets
}
