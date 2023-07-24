package main

import (
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
	r.GET("/profile", ProfileController)
	r.GET("/change/:id", ChangeController)
	r.POST("/create", CreatePostController)
	r.POST("/login", LoginUserPostController)
	r.POST("/register", RegisterPostUser)
	r.POST("/change", ChangePostController)
	r.GET("/api/findArts", FindArts)
	r.DELETE("/api/deleteArt", ProfileDeleteArt)
	//r.PUT("/changeUser", ChangeUser)
	r.Run(":8080")
}
