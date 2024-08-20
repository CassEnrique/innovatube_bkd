/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package utils

import (
	"log"
	"time"

	"github.com/backend/model"
	"github.com/golang-jwt/jwt/v5"
)

func JwtLoginUser(login model.JwtLoginUser) string {
	claims := jwt.MapClaims{
		"user":  login.User,
		"email": login.Email,
		"pk":    login.Pk,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(JwtSecretKey()))
	if err != nil {
		log.Printf("Fallo al crear el token JWT token.SignedString() en la funcion jwt.JwtLoginCurp()")
		return ""
	}

	return t
}
