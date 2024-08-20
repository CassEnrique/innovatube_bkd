/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package users

import (
	"log"

	"github.com/backend/model"
	"github.com/backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx, db *gorm.DB) interface{} {
	user := model.User{}

	if err := c.BodyParser(&user); err != nil {
		log.Print("Fallo al construir el modelo model.User{}, en la funcion user.Create()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword
	userCreated := db.Create(&user)

	if userCreated.Error != nil {
		return fiber.Map{
			"code": 204,
			"message": []string{
				"Sin datos.",
			},
			"err": userCreated.Error,
		}
	}

	user.Password = ""
	return fiber.Map{
		"code": 200,
		"message": []string{
			"Usuario se ha creado exitosamente!!!.",
		},
		"obj": user,
	}
}

func Get(pk string, db *gorm.DB) interface{} {
	user := model.User{}

	db.Find(&user, pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Usuario se ha obtenido exitosamente!!!.",
		},
		"obj": user,
	}
}

func Update(c *fiber.Ctx, db *gorm.DB) interface{} {
	patching := model.Patching{}
	user := model.User{}

	if err := c.BodyParser(&user); err != nil {
		log.Print("Fallo al construir el modelo model.{}, en la funcion .()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	db.Model(&user).Where("pk_user = ?", patching.Pk).Update(patching.Key, patching.Value)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Usuario se ha actualizado exitosamente!!!.",
		},
		"obj": user,
	}
}

func Delete(pk string, db *gorm.DB) interface{} {
	user := model.User{}

	db.Delete(&user, pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Usuario se ha eliminado exitosamente!!!.",
		},
		"obj": user,
	}
}
