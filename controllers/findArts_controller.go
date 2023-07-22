package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
	"fmt"
	"strings"
)

func FindArts(c *gin.Context) {
	var userIsLogin, _ = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	findValue := c.Query("find")

	var arts []Art

	if findValue != "" {
		findValue = "%" + strings.ToLower(findValue) + "%"
		DB.Where("lower(Name) LIKE ? OR lower(Description) LIKE ?", findValue, findValue).Find(&arts)
	} else {
		DB.Find(&arts)
	}

  var simplifiedArts []map[string]interface{}
	for _, art := range arts {
		simplifiedArt := map[string]interface{}{
			"Name":        art.Name,
			"Image":       art.Image,
			"Description": art.Description,
			"Price":       art.Price,
		}
		simplifiedArts = append(simplifiedArts, simplifiedArt)
	}

	c.JSON(http.StatusOK, simplifiedArts)
}

