/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package login

import (
	"log"

	"github.com/backend/model"
	"github.com/backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, db *gorm.DB) interface{} {
	user := model.User{}
	login := model.Login{}
	jwtLogin := model.JwtLoginUser{}

	if err := c.BodyParser(&login); err != nil {
		log.Print("Fallo al construir el modelo model.Login{}, en la funcion login.Login()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	userSelected := db.Where("username = ?", login.User).Or("email = ?", login.User).Find(&user)

	if userSelected.Error != nil {
		return fiber.Map{
			"code": 204,
			"message": []string{
				"Sin datos.",
			},
			"err": userSelected.Error,
		}
	}

	checkOassword := utils.CheckPasswordHash(login.Password, user.Password)

	if !checkOassword {
		return fiber.Map{
			"code": 203,
			"message": []string{
				"Contraseña no valida.",
			},
			"err": nil,
		}
	}

	jwtLogin.Pk = user.Pk
	jwtLogin.Email = user.Email
	jwtLogin.User = user.Username
	token := utils.JwtLoginUser(jwtLogin)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Usuario obtenido exitosamente!!!.",
		},
		"obj":   user,
		"token": token,
	}
}
