/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package login

import (
	"github.com/backend/crud/login"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handlerLogin(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(login.Login(c, db))
}
