package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func HomeController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	var arts []Art
	DB.Find(&arts)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Kris Arts",
		"arts":        arts,
		"user":        user,
		"userIsLogin": userIsLogin,
	})
}
