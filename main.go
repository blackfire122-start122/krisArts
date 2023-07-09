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
	r.POST("/login", LoginUserPostController)
	r.GET("/logout", LogoutUser)
	r.GET("/register", RegisterUser)
	r.POST("/register", RegisterPostUser)
	//r.PUT("/changeUser", ChangeUser)
	r.POST("/create", CreatePostController)
	r.Run(":8080")
}