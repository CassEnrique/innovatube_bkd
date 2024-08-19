/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package favorites

import (
	"github.com/backend/crud/favorites"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handlerCreateFavorite(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(favorites.Create(c, db))
}

func handlerGetFavorite(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(favorites.Get(pk, db))
}

func handlerGetAllFavorite(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(favorites.GetAll(c, db))
}

func handlerUpdateFavorite(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(favorites.Update(c, db))
}

func handlerDeleteFavorite(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(favorites.Delete(pk, db))
}
