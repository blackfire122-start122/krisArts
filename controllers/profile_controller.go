package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func ProfileController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	var arts []Art
	DB.Find(&arts, "user_id=?",user.Id).Limit(20)

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title":       "Profile",
		"arts":        arts,
		"user":        user,
		"userIsLogin": userIsLogin,
	})
}
