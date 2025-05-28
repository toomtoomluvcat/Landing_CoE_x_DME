package service

import (
	"fmt"
	"scavenger_hunt/app/connect"
	"scavenger_hunt/app/model"

	"github.com/gin-gonic/gin"
)

func Create_Account(c *gin.Context) {
	var req struct {
		email       string
		user_name   string
		description string
		link_1      string
		link_2      string
		link_3      string
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := connect.DB.Model(&model.Participant{}).Where("email = ?", req.email).
		Updates(map[string]interface{}{
			"user_name":   req.user_name,
			"description": req.description,
			"link_1":      req.link_1,
			"link_2":      req.link_2,
			"link_3":      req.link_3}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Fail to update"})
		return
	}
	c.Status(201)
}

func Check_Account(c *gin.Context) {
	var email string
	if err := connect.DB.Model(&model.Participant{}).Select("email").First(&email); err != nil {
		c.JSON(404, gin.H{"error": "email not found"})
		return
	}

}

func Join_Account(c *gin.Context) {
	var req model.Join
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid json payload"})
	}

	var count int64
	if err := connect.DB.Model(&model.Join{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		c.JSON(500, gin.H{"error": "fail to seacrh email"})
		return
	}

	if count < 1 {
		if err := connect.DB.Model(&model.Join{}).Create(&req).Error; err != nil {
			fmt.Printf("error at :%v", err)
			c.JSON(500, gin.H{"error": fmt.Sprintf("Fail to create new join: %v", err)})
			return
		}
	} else {
		if err := connect.DB.Model(&model.Join{}).Where("email =?", req.Email).Updates(map[string]interface{}{"profile": req.Profile}).Error; err != nil {
			fmt.Printf("error at :%v", err)
			c.JSON(500, gin.H{"error": "Fail to update"})
		}
	}

	fmt.Printf("email : %s \nprofile: %s\n", req.Email, req.Profile)
	c.JSON(201, gin.H{"message": "user join successfuly"})
}

func Get_Account(c *gin.Context) {
	var response []string
	if err := connect.DB.Model(&model.Join{}).Order("random()").Limit(5).Select("profile").Find(&response).Error; err != nil {
		c.JSON(500, gin.H{"error": "fail to insert"})
		return
	}
	var count int64
	if err := connect.DB.Model(&model.Join{}).Select("profile").Find(&response).Count(&count).Error; err != nil {
		c.JSON(500, gin.H{"error": "fail to insert"})
		return
	}
	c.JSON(200, gin.H{"avatar": response, "count": count})
}
