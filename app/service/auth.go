package service

import (
	"fmt"
	"scavenger_hunt/app/connect"
	"scavenger_hunt/app/model"

	"github.com/gin-gonic/gin"
)

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

	var count int64
	if err := connect.DB.Model(&model.Join{}).Select("profile").Find(&response).Count(&count).Error; err != nil {
		c.JSON(500, gin.H{"error": "fail to insert"})
		return
	}
	if err := connect.DB.
		Model(&model.Join{}).Limit(5).
		Order("random()").
		Pluck("profile", &response).Error; err != nil {
		c.JSON(500, gin.H{"error": "fail to fetch"})
		return
	}
	c.JSON(200, gin.H{"avatar": response, "count": count})
}
