package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func OrderController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	if err := DB.Preload("Basket.Arts").First(&user).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var fullPrice float64 = 0

	for _, art := range user.Basket.Arts {
		fullPrice += art.Price
	}

	var order Order

	order.Price = fullPrice
	order.User = user
	order.Arts = user.Basket.Arts

	if err := DB.Create(&order).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := DB.Model(&user.Basket).Association("Arts").Clear(); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"price": fullPrice,
		"arts":  order.Arts,
	})
}

func FindCity(c *gin.Context) {
	searchWord := c.PostForm("searchWord")

	resp, err := FindCityApiNovaPoshta(searchWord)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, resp)
}

func GetWarehouses(c *gin.Context) {
	cityRef := c.PostForm("cityRef")

	resp, err := GetWarehousesApiNovaPoshta(cityRef)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, resp)
}

func SearchSettlements(c *gin.Context) {
	cityRef := c.PostForm("cityRef")

	resp, err := GetSettlementsApiNovaPoshta(cityRef)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, resp)
}
