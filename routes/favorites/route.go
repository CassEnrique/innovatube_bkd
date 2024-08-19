/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package favorites

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitFavoriteRoutes(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("/v1")
	favoriteV1 := v1.Group("/favorite")

	favoriteV1.Post("/create", func(c *fiber.Ctx) error {
		return handlerCreateFavorite(c, db)
	})

	favoriteV1.Get("/get/pk/:pk", func(c *fiber.Ctx) error {
		return handlerGetFavorite(c, db)
	})

	favoriteV1.Get("/get/all", func(c *fiber.Ctx) error {
		return handlerGetAllFavorite(c, db)
	})

	favoriteV1.Patch("/patch", func(c *fiber.Ctx) error {
		return handlerUpdateFavorite(c, db)
	})

	favoriteV1.Delete("/delete/:pk", func(c *fiber.Ctx) error {
		return handlerDeleteFavorite(c, db)
	})
}
