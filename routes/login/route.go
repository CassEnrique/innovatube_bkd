/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package login

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitLoginRoutes(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("/v1")
	loginV1 := v1.Group("/login")

	loginV1.Post("/user", func(c *fiber.Ctx) error {
		return handlerLogin(c, db)
	})
}
