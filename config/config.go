package config

import (
	"fmt"
	"os"
)

func JwtSecretKey() string {
	key := fmt.Sprintf("%+v-%+v", os.Getenv("JWT_SECRET_KEY"), os.Getenv("APP"))
	//secret, err := utils.HashPassword(key)

	//if err != nil {
	//	log.Printf("Fallo al construir el utils.HashPassword() en la funcion JwtSecretKey(): %+v", err)
	//}

	return key
}
