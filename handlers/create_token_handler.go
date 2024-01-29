package handlers

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("your-secret-key") //change this to your secret key

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
