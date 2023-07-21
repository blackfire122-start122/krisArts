package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func FindArts(c *gin.Context) {
	var userIsLogin, _ = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	var arts []Art

	DB.Find(&arts, "")

	c.JSON(http.StatusOK, nil)
}
