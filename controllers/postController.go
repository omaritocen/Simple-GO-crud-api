package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	// Return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	// Return it
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	// Return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(http.StatusNoContent)
}

