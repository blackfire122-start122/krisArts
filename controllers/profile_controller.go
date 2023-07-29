package controllers

import (
	"github.com/gin-gonic/gin"
	. "krisArts/models"
	. "krisArts/utils"
	"net/http"
	"os"
	"strconv"
)

func ProfileController(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)
	
	var userOfParam User
	username := c.Param("username")
	
	if err:= DB.First(&userOfParam,"username=?", username).Error; err!=nil {
	  c.Writer.WriteHeader(http.StatusBadRequest)
	  return
	}

	var arts []Art
	DB.Find(&arts, "user_id=?", userOfParam.Id).Limit(20)

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title":       "Profile",
		"arts":        arts,
		"user":        user,
		"userIsLogin": userIsLogin,
		"userProfile": userOfParam,
		"profileOwner":user.Id==userOfParam.Id,
	})
}

func ProfileDeleteArt(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	deleteId := c.Query("id")

	var art Art

	if err := DB.First(&art, deleteId).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	if art.UserId != user.Id {
	  c.Writer.WriteHeader(http.StatusForbidden)
	  return
	}

	if art.Image != "" {
		if _, err := os.Stat(art.Image); err == nil {
			if err := os.Remove(art.Image); err != nil {
				c.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else if os.IsNotExist(err) {

		} else {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if err := DB.Delete(Art{}, deleteId).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func LoadArtsUser(c *gin.Context) {
	var userIsLogin, user = CheckSessionUser(c.Request)

	if !userIsLogin {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}

	countArts, err := strconv.Atoi(c.Query("countArts"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var arts []Art

	if err := DB.Limit(20).Offset(countArts).Find(&arts, "user_id=?", user.Id).Error; err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var simplifiedArts []map[string]interface{}
	for _, art := range arts {
		simplifiedArt := map[string]interface{}{
			"Name":        art.Name,
			"ID":          art.ID,
			"Image":       art.Image,
			"Description": art.Description,
			"Price":       art.Price,
		}
		simplifiedArts = append(simplifiedArts, simplifiedArt)
	}

	c.JSON(http.StatusOK, simplifiedArts)
}
