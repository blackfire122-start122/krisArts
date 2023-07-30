package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	"net/http"
	"strconv"
	"strings"
)

func FindArts(c *gin.Context) {
	countArts, err := strconv.Atoi(c.Query("countArts"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	findValue := c.Query("find")

	var arts []Art

	if findValue != "" {
		findValue = "%" + strings.ToLower(findValue) + "%"
		DB.Where("lower(Name) LIKE ? OR lower(Description) LIKE ?", findValue, findValue).Limit(20).Offset(countArts).Find(&arts)
	} else {
		DB.Limit(20).Offset(countArts).Find(&arts)
	}

	var simplifiedArts []map[string]interface{}
	for _, art := range arts {
		simplifiedArt := map[string]interface{}{
			"Name":        art.Name,
			"Image":       art.Image,
			"Description": art.Description,
			"Price":       art.Price,
			"ID":          art.ID,
		}
		simplifiedArts = append(simplifiedArts, simplifiedArt)
	}

	c.JSON(http.StatusOK, simplifiedArts)
}
