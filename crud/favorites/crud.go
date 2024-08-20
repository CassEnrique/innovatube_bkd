/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package favorites

import (
	"log"

	"github.com/backend/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx, db *gorm.DB) interface{} {
	favorite := model.Favorite{}

	if err := c.BodyParser(&favorite); err != nil {
		log.Print("Fallo al construir el modelo model.Favorite{}, en la funcion favorites.Create()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	favoriteExists := model.Favorite{}
	if favoriteFind := db.Find(&favoriteExists, "video_playlist_id = ?", favorite.VideoPlaylistId); favoriteFind.RowsAffected == 1 {
		return fiber.Map{
			"code": 204,
			"message": []string{
				"Este video ya esta en tus favoritos!!!.",
			},
			"obj": model.Favorite{},
		}
	}

	favoriteCreated := db.Create(&favorite)

	if favoriteCreated.Error != nil {
		return fiber.Map{
			"code": 204,
			"message": []string{
				"Sin datos.",
			},
			"err": favoriteCreated.Error,
		}
	}

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha creado exitosamente!!!.",
		},
		"obj": favorite,
	}
}

func Get(pk string, db *gorm.DB) interface{} {
	favorite := model.Favorite{}

	db.Find(&favorite, pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Favorito se ha obtenido exitosamente!!!.",
		},
		"obj": favorite,
	}
}

func GetAll(c *fiber.Ctx, db *gorm.DB) interface{} {
	favorites := model.Favorites{}

	db.Find(&favorites)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Favoritos se ha obtenido exitosamente!!!.",
		},
		"obj": favorites,
	}
}

func Update(c *fiber.Ctx, db *gorm.DB) interface{} {
	patching := model.Patching{}
	favorite := model.Favorite{}

	if err := c.BodyParser(&favorite); err != nil {
		log.Print("Fallo al construir el modelo model.{}, en la funcion .()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	db.Model(&favorite).Where("pk_favorite = ?", patching.Pk).Update(patching.Key, patching.Value)

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha actualizado exitosamente!!!.",
		},
		"obj": favorite,
	}
}

func Delete(pk string, db *gorm.DB) interface{} {
	favorite := model.Favorite{}

	db.Delete(&favorite, "video_playlist_id = ?", pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			"Video ya no es tu favorito!!!.",
		},
		"obj": favorite,
	}
}
