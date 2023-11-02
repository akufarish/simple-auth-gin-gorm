package authController

import (
	"authGin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Payload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context)  {
	var user models.User

	c.BindJSON(&user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
	}

	user.Password = string(hash)

	models.DB.Create(&user)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Register successfull",
		"user": user,
	})
}

func Login (c *gin.Context) {
	var user models.User
	var payload Payload

	c.BindJSON(&payload)

	models.DB.Where("email = ?", payload.Email).First(&user)
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "incorrect password",
		})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Login successfull",
		"user": user,
	})
}