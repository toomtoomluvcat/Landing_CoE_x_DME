package service

import (
	"errors"
	"fmt"
	"scavenger_hunt/app/connect"
	"scavenger_hunt/app/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTeam(c *gin.Context) {
	var req struct {
		Name string
		Code string
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	group := model.Group{
		Name: req.Name,
		Code: req.Code,
	}

	if err := connect.DB.Create(&group).Error; err != nil {
		c.JSON(500, gin.H{"error": "Fail to create Group" + err.Error()})
		fmt.Println(err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "sucess fully to create"})
}

func JoinTeam(c *gin.Context) {
	var req struct {
		Email string
		Code  string
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": "Invalid JSON payload"})
		return
	}

	var group model.Group
	if err := connect.DB.Model(&model.Group{}).Where("Code = ?", req.Code).First(&group).Error; err != nil {
		c.JSON(404, gin.H{"error": "Invalid Team code"})
		return
	}
	err := connect.DB.Transaction(func(tx *gorm.DB) error {
		var join model.Join
		if err := tx.Where("email = ?", req.Email).First(&join).Error; err != nil {
			return errors.New("user not register")
		}

		if join.GroupID != nil && *join.GroupID == group.ID {
			return errors.New("you already joined this team")
		}

		join.GroupID = &group.ID
		if err := tx.Save(&join).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "successfully to join group", "ID": group.ID})
}

func GetAllTeam(c *gin.Context) {
	var groups []model.Group

	if err := connect.DB.Model(&model.Group{}).Preload("Joins").Find(&groups).Error; err != nil {
		c.JSON(500, gin.H{"error": "Fail to get group"})
		return
	}
	c.JSON(200, groups)
}

func GetProblem(c *gin.Context) {
	var req struct {
		Code string
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	switch req.Code {
	case "python":
		c.JSON(200, gin.H{"message": "lkhs5"})
		return
	case "java":
		c.JSON(200, gin.H{"message": "lkshs6"})
		return
	default:
		c.JSON(400, gin.H{"error": "incorrect answer try again"})
		return
	}
}
