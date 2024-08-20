/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitUserRoutes(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("/v1")
	userV1 := v1.Group("/user")
	
	userV1.Post("/create", func(c *fiber.Ctx) error {
		return handlerCreateUser(c, db)
	})

	userV1.Get("/get/pk/:pk", func(c *fiber.Ctx) error {
		return handlerGetUser(c, db)
	})

	userV1.Patch("/patch", func(c *fiber.Ctx) error {
		return handlerUpdateUser(c, db)
	})

	userV1.Delete("/delete/:pk", func(c *fiber.Ctx) error {
		return handlerDeleteUser(c, db)
	})
}
