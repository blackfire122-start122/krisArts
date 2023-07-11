package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

type Art struct {
	ID          uint   `gorm:"primaryKey"`
	Image       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       float64
}

func main() {
	r := gin.Default()

	db := initDB()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/media/", "./media")

	r.GET("/", func(c *gin.Context) {
		var arts []Art
		db.Find(&arts)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"arts": arts,
		})
	})

	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})

	r.POST("/create", func(c *gin.Context) {
		image, _ := c.FormFile("image")
		description := c.PostForm("description")
		price := c.PostForm("price")

		// Збереження зображення у папці media/arts
		imagePath := filepath.Join("media", "arts", image.Filename)
		if err := saveImage(image, imagePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Помилка при збереженні файлу: %s", err.Error()))
			return
		}

		art := Art{
			Image:       imagePath,
			Description: description,
			Price:       parsePrice(price),
		}
		db.Create(&art)

		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Не вдалося підключитись до бази даних")
	}

	db.AutoMigrate(&Art{})

	return db
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
	// Тут можна додати логіку для перевірки та обробки ціни
	// Наприклад, перевірка на коректність числа та форматування
	// У цьому прикладі просто конвертуємо рядок в float64
	price := 0.0
	fmt.Sscanf(priceStr, "%f", &price)
	return price
}
