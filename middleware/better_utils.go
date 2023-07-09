package middleware

import (
	"Cursorr/BankingSystem/database"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}
		return []byte(os.Getenv("SECRET")), nil
	})
}

func getUserFromToken(tokenString string) (database.User, error) {
	token, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return database.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
		return database.User{}, fmt.Errorf("invalid token")
	}

	email, ok := claims["sub"].(string)
	if !ok {
		return database.User{}, fmt.Errorf("invalid token")
	}

	user := database.Instance.GetUserByEmail(email)
	if user == (database.User{}) {
		return database.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}
