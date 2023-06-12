package controllers

import (
	"net/http"
	"zappin/models"
	"zappin/services"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Email string `json:"email" binding:"required"`
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
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	user.Email = data.Email
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
