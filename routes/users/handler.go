/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package users

import (
	"github.com/backend/crud/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handlerCreateUser(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(users.Create(c, db))
}

func handlerGetUser(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(users.Get(pk, db))
}

func handlerUpdateUser(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(users.Update(c, db))
}

func handlerDeleteUser(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(users.Delete(pk, db))
}
