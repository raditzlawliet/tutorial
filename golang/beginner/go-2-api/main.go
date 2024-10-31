package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create new GET route /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// groping route with /arith
	arith := r.Group("/arith")
	{
		// define GEt route /arith/sum
		arith.GET("/sum", func(c *gin.Context) {
			// Query String and parse to int to variable
			a, _ := strconv.Atoi(c.DefaultQuery("a", "0"))
			b, _ := strconv.Atoi(c.DefaultQuery("b", "0"))

			// Response the result
			c.JSON(http.StatusOK, gin.H{
				"message": sum(a, b),
			})
		})

		// defined POST route /arith/sub
		arith.POST("/sub", arithSubHandler)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type arithSubData struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func arithSubHandler(c *gin.Context) {
	var data arithSubData

	// Bind the JSON payload to the User struct
	if err := c.ShouldBindJSON(&data); err != nil {
		// Return a 400 Bad Request response with the error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Response the result
	c.JSON(http.StatusOK, gin.H{
		"message": sub(data.A, data.B),
	})
}

// simple arith
func sum(a int, b int) int {
	return a + b
}

// simple arith
func sub(a int, b int) int {
	return a - b
}
