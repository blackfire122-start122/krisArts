package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
	"os"
)

func ChangeController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	var art Art
	id := c.Param("id")

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

	id := c.Param("id")

	var art Art
	if err := DB.First(&art, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при знаходженні арту: %s", err.Error()))
		return
	}

	image, err := c.FormFile("image")
	if err == nil {
		if art.Image != "" {
			if _, err := os.Stat(art.Image); err == nil {
				if err := os.Remove(art.Image); err != nil {
					c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при видаленні старої картинки: %s", err.Error()))
					return
				}
			} else if os.IsNotExist(err) {

			} else {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при перевірці файлу: %s", err.Error()))
				return
			}
		}

		imagePath := "media/arts/" + id + image.Filename

		if err := saveImage(image, imagePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при збереженні нової картинки: %s", err.Error()))
			return
		}

		art.Image = imagePath
	}

	description := c.PostForm("description")
	price := c.PostForm("price")
	name := c.PostForm("name")

	art.Description = description
	art.Price = parsePrice(price)
	art.Name = name

	if err := DB.Save(&art).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при оновленні арту: %s", err.Error()))
		return
	}

	c.Redirect(http.StatusSeeOther, "/profile")
}
