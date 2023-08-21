package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
	"strconv"
)

func AddToBasket(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	artIDStr := c.PostForm("artId")
	artID, err := strconv.ParseUint(artIDStr, 10, 64)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	DB.Preload("Basket").First(&user)

	var artToAdd Art
	if err := DB.First(&artToAdd, artID).Error; err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Basket.Arts = append(user.Basket.Arts, artToAdd)

	if err := DB.Save(&user).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func GetAllArtsBasket(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	DB.Preload("Basket.Arts").First(&user)

	var resp []map[string]interface{}
	for _, art := range user.Basket.Arts {
		artItem := map[string]interface{}{
			"Name":        art.Name,
			"Image":       art.Image,
			"Description": art.Description,
			"Price":       art.Price,
			"ID":          art.ID,
		}
		resp = append(resp, artItem)
	}

	c.JSON(http.StatusOK, resp)

}
