package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/utils"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Create User",
	})
}

func RegisterPostUser(c *gin.Context) {
	var form UserRegister
	if err := c.ShouldBind(&form); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if form.Password == "" || form.Username == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := Sign(&form, c); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func LoginUserController(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Enter",
	})
}

func LoginUserPostController(c *gin.Context) {
	var form UserLogin

	if err := c.ShouldBind(&form); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if Login(c.Writer, c.Request, &form) {
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": "error"})
	}
}

func LogoutUser(c *gin.Context) {
	if Logout(c.Writer, c.Request) {
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
}
