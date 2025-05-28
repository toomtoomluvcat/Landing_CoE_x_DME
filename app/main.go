package main

import (
	"scavenger_hunt/app/connect"
	"scavenger_hunt/app/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Pong"})
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	connect.Pg()
	router.POST("/user", service.Check_Account)
	router.POST("/join", service.Join_Account)
	router.GET("/profile", service.Get_Account)
	router.GET("/Ping", Ping)
	router.Run(":8000")
}
