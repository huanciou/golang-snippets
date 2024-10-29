package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password"`
}

func main() {

	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "200",
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err == nil {
			fmt.Println(user)

			c.JSON(200, user)
		} else {
			fmt.Println(user)

			c.JSON(404, gin.H{
				"Status": "Validation Error",
			})
		}
	})

	r.GET("redirect", func(c *gin.Context) {
		c.Redirect(307, "https://google.com")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"result": "404 Not Found",
		})
	})

	r.Run(":3000")
}
