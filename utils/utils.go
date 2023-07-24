package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	. "krisArts/models"
	"mime/multipart"
	"net/http"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))

type UserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(w http.ResponseWriter, r *http.Request, userLogin *UserLogin) bool {
	session, _ := store.Get(r, "session-name")

	var user User
	if err := DB.First(&user, "Username = ?", userLogin.Username).Error; err != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err == nil {
		session.Values["id"] = user.Id
		session.Values["password"] = user.Password
		err = session.Save(r, w)
		if err != nil {
			return false
		}
	} else {
		return false
	}

	return true
}

type UserRegister struct {
	Username string                `form:"username" binding:"required"`
	Password string                `form:"password" binding:"required"`
	Image    *multipart.FileHeader `form:"image"    binding:"required"`
}

func Sign(user *UserRegister, c *gin.Context) error {
	var users []User

	if err := DB.Where("Username = ?", user.Username).Find(&users).Error; err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("user with the same username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	var ImageName string

	if user.Image.Filename != "" && user.Image.Size != 0 {
		if err := c.SaveUploadedFile(user.Image, "./media/UserImages/"+user.Username+user.Image.Filename); err != nil {
			return err
		}
		ImageName = "/media/UserImages/" + user.Username + user.Image.Filename
	}

	err = DB.Create(&User{Username: user.Username, Password: string(hashedPassword), Image: ImageName}).Error
	return err
}

func CheckSessionUser(r *http.Request) (bool, User) {
	session, _ := store.Get(r, "session-name")

	var user User

	if session.IsNew {
		return false, user
	}

	if err := DB.First(&user, "Id = ?", session.Values["id"]).Error; err != nil {
		return false, user
	}

	if session.Values["password"] != user.Password {
		return false, user
	}
	return true, user
}

func CheckAdmin(user User) bool {
	var admin Admin
	if err := DB.Where("user_id=?", user.Id).Find(&admin).Error; err != nil {
		return false
	}

	return admin.UserId == user.Id
}

func Logout(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session-name")

	session.Values["id"] = nil
	session.Values["password"] = nil

	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		return false
	}
	return true
}
