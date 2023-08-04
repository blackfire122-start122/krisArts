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
	if err := DB.Limit(60).Find(&arts).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	var artsLines []Art

	if len(arts) > 0 {
		for len(artsLines) < 60 {
			if len(artsLines)+len(arts) <= 60 {
				artsLines = append(artsLines, arts...)
			} else {
				remainingSpace := 60 - len(artsLines)
				artsLines = append(artsLines, arts[:remainingSpace]...)
			}
		}
	} else {
		artsLines = make([]Art, 60)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Kris Arts",
		"arts":        arts,
		"user":        user,
		"userIsLogin": userIsLogin,
		"artsLine1":   artsLines[:20],
		"artsLine2":   artsLines[20:40],
		"artsLine3":   artsLines[40:],
	})
}
