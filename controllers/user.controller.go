package controllers

import (
	"fmt"
	"net/http"
	"zappin/models"
	"zappin/services"

	"github.com/gin-gonic/gin"
)

type UserSignUpData struct {
	Email string `json:"email" binding:"required"`
}

type UserData struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

//TODO: remove this endpoint before deployment to prduction

func AllUsers(c *gin.Context) {
	var users []models.User
	if err := services.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

func AddUser(c *gin.Context) {
	var data UserSignUpData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	user.Email = data.Email
	user.Musics = []models.Post{}
	if err := services.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func UpdateUser(c *gin.Context) {
	var data UserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	fmt.Println(data.ID)
	if err := services.DB.Find(&user).Where("id = ?", data.ID).Update("name", data.Name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	services.DB.Save(&user)
	//services.DB.Commit()
	//user.Name = data.Name

	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   user,
	})
}
