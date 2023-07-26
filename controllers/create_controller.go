package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	. "krisArts/models"
	. "krisArts/utils"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func CreateController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)
	c.HTML(http.StatusOK, "create.html", gin.H{
		"title":       "Create Art",
		"user":        user,
		"userIsLogin": userIsLogin,
	})
}

func CreatePostController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	image, _ := c.FormFile("image")
	description := c.PostForm("description")
	price := c.PostForm("price")
	name := c.PostForm("name")

	art := Art{
		Description: description,
		Price:       parsePrice(price),
		User:        user,
		Name:        name,
	}

	if err := DB.Create(&art).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	imagePath := "media/arts/" + strconv.Itoa(int(art.ID)) + image.Filename

	if err := saveImage(image, imagePath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при збереженні файлу: %s", err.Error()))
		return
	}

	art.Image = imagePath
	if err := DB.Save(&art).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func saveImage(file *multipart.FileHeader, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func parsePrice(priceStr string) float64 {
	price := 0.0
	fmt.Sscanf(priceStr, "%f", &price)
	return price
}
