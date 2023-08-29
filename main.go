package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "krisArts/controllers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/media/", "./media")
	r.Static("/static/", "./static")

	r.GET("/", HomeController)
	r.GET("/create", CreateController)
	r.GET("/login", LoginUserController)
	r.GET("/logout", LogoutUser)
	r.GET("/register", RegisterUser)
	r.GET("/profile/:username", ProfileController)
	r.GET("/change/:id", ChangeController)
	r.POST("/create", CreatePostController)
	r.POST("/login", LoginUserPostController)
	r.POST("/register", RegisterPostUser)
	r.POST("/change/:id", ChangePostController)
	r.GET("/api/findArts", FindArts)
	r.DELETE("/api/deleteArt", ProfileDeleteArt)
	r.GET("/api/profile/loadArtsUser", LoadArtsUser)
	r.GET("art/:id", ArtController)
	r.POST("/api/addToBasket", AddToBasket)
	r.GET("/api/getAllArtsBasket", GetAllArtsBasket)
	r.DELETE("/api/deleteFromBasket/:artId", DeleteFromBasket)
	r.POST("/api/order", OrderController)
	r.POST("/api/order/findCityNovaPoshta", FindCity)
	r.POST("/api/order/getWarehouses", GetWarehouses)
	r.POST("/api/order/GetSettlements", SearchSettlements)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Err run server ", err.Error())
		return
	}

}
