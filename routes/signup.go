package routes

import (
	"Cursorr/BankingSystem/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignupRoute() gin.HandlerFunc {
	return Signup
}

func Signup(c *gin.Context) {
	var body struct {
		Email     string
		Password  string
		Firstname string
		Lastname  string
		Age       int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := database.User{
		FirstName: body.Firstname,
		LastName:  body.Lastname,
		Email:     body.Email,
		Age:       body.Age,
		Password:  string(hash),
	}

	err = database.Instance.CrateNewUser(user)

	println(err.Error())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
