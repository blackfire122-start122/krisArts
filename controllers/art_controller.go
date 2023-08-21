package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func ArtController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)
	
	id := c.Param("id")

	var art Art
	DB.Preload("User").First(&art, "id=?", id)

	c.HTML(http.StatusOK, "art.html", gin.H{
		"title":       "Art",
		"art":         art,
		"user":        user,
		"userIsLogin": userIsLogin,
	})
}