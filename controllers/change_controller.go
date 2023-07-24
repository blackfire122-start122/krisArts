package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
)

func ChangeController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	var art Art
	id := c.Query("id")

	if err := DB.First(&art, id).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if art.UserId != user.Id {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	c.HTML(http.StatusOK, "change.html", gin.H{
		"title":       "Change Art",
		"user":        user,
		"userIsLogin": userIsLogin,
		"art":         art,
	})
}

func ChangePostController(c *gin.Context) {
	var userIsLogin, _ = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	//image, _ := c.FormFile("image")
	//description := c.PostForm("description")
	//price := c.PostForm("price")
	//name := c.PostForm("name")
	//
	//imagePath := "media/arts/" + image.Filename
	//if err := saveImage(image, imagePath); err != nil {
	//	c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при збереженні файлу: %s", err.Error()))
	//	return
	//}

	//art := Art{
	//	Image:       imagePath,
	//	Description: description,
	//	Price:       parsePrice(price),
	//	User:        user,
	//	Name:        name,
	//}
	//DB.Create(&art)

	c.Redirect(http.StatusSeeOther, "/profile")
}
