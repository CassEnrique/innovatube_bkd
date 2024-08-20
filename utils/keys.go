/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package utils

import (
	"fmt"
	"os"
)

func JwtSecretKey() string {
	key := fmt.Sprintf("%+v-%+v", os.Getenv("JWT_SECRET_KEY"), os.Getenv("APP"))

	return key
}
