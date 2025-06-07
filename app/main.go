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
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	connect.Pg()
	route.POST("/join", service.Join_Account)
	route.GET("/profile", service.Get_Account)
	route.POST("/profile", service.GetProfile)
	route.PUT("/profile", service.SetProfile)
	route.POST("/team", service.CreateTeam)
	route.PUT("/team", service.JoinTeam)
	route.GET("/team", service.GetAllTeam)
	route.POST("/problem", service.GetProblem)
	route.GET("/Ping", Ping)
	route.Run()
}
