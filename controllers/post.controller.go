package controllers

import (
	"net/http"
	"zappin/models"
	"zappin/services"

	"github.com/gin-gonic/gin"
)

type Postdata struct {
	AuthorID    uint   `json:"author_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"desc"`
	Genre       string `json:"genre" binding:"required"`
	Label       string `json:"label" binding:"required"`
}

type MusicData struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func AddMusic(c *gin.Context) {
	var musicData MusicData
	if err := c.ShouldBindJSON(&musicData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := models.Post{}
	if err := services.DB.Find(&post).Where("id = ?", musicData.ID).Update("music_url", musicData.Url).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
	})

}

func AddPost(c *gin.Context) {
	var postData Postdata
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	post := models.Post{}
	post.AuthorID = postData.AuthorID
	post.Name = postData.Name
	post.Description = postData.Description
	post.Genre = postData.Genre
	post.Label = postData.Label
	if err := services.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   post,
	})
}

func AllPosts(c *gin.Context) {
	var posts []models.Post
	if err := services.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   posts,
	})
}
