package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
	"os"
)

func ProfileController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	var arts []Art
	DB.Find(&arts, "user_id=?", user.Id).Limit(20)

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title":       "Profile",
		"arts":        arts,
		"user":        user,
		"userIsLogin": userIsLogin,
	})
}

func ProfileDeleteArt(c *gin.Context) {
	var userIsLogin, _ = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	deleteId := c.Query("id")

	var art Art

	if err := DB.First(&art, deleteId).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if art.Image != "" {
		if _, err := os.Stat(art.Image); err == nil {
			if err := os.Remove(art.Image); err != nil {
				c.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else if os.IsNotExist(err) {

		} else {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if err := DB.Delete(Art{}, deleteId).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
