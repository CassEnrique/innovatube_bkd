#!/bin/bash

file_name=$1

author='eecass'
datenow= date +"%m/&d/%Y %H:%M"

header="/*
@Author: $author
@Date:   $datenow
@Last Modified by:   $author
@Last Modified time: $datenow
*/"

initCrud='
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func Create(c *fiber.Ctx, db *gorm.DB) interface{} {
	 := model.{}

	if err := c.BodyParser(&); err != nil {
		log.Print("Fallo al construir el modelo model.{}, en la funcion .()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	Created := db.Create(&)

	if .Error != nil {
		return fiber.Map{
			"code": 204,
			"message": []string{
				"Sin datos.",
			},
			"err": .Error,
		}
	}

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha creado exitosamente!!!.",
		},
		"obj": ,
	}
}

func Get(pk string, db *gorm.DB) interface{} {
	 := model.{}

	db.Find(&, pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha obtenido exitosamente!!!.",
		},
		"obj": ,
	}
}

func Update(c *fiber.Ctx, db *gorm.DB) interface{} {
	patching := model.Patching{}
	 := model.{}

	if err := c.BodyParser(&); err != nil {
		log.Print("Fallo al construir el modelo model.{}, en la funcion .()")

		return fiber.Map{
			"code": 500,
			"message": []string{
				"Error en el servidor.",
			},
			"err": err,
		}
	}

	db.Model(&).Where("pk_ = ?", patching.Pk).Update(patching.Key, patching.Value)

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha actualizado exitosamente!!!.",
		},
		"obj": ,
	}
}

func Delete(pk string, db *gorm.DB) interface{} {
	 := model.{}

	db.Delete(&, pk)

	return fiber.Map{
		"code": 200,
		"message": []string{
			" se ha eliminado exitosamente!!!.",
		},
		"obj": ,
	}
}'

initHandler='
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handlerCreate_(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(_.Create(c, db))
}

func handlerGet_(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(_.Get(pk, db))
}

func handlerUpdate_(c *fiber.Ctx, db *gorm.DB) error {
	return c.JSON(_.Update(c, db))
}

func handlerDelete_(c *fiber.Ctx, db *gorm.DB) error {
	pk := c.Params("pk")

	return c.JSON(_.Delete(pk, db))
}'

initRoute='
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init_Routes(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("/v1")
	_V1 := v1.Group("/_")

	_V1.Post("/create", func(c *fiber.Ctx) error {
		return handlerCreate_(c, db)
	})

	_V1.Get("/get/:pk", func(c *fiber.Ctx) error {
		return handlerGet_(c, db)
	})

	_V1.Patch("/patch", func(c *fiber.Ctx) error {
		return handlerUpdate_(c, db)
	})

	_V1.Delete("/delete/:pk", func(c *fiber.Ctx) error {
		return handlerDelete_(c, db)
	})
}'

function CreateLayout() {
    mkdir crud/$file_name
    echo -e "$header\n\npackage $file_name\n\n$initCrud" > crud/$file_name/crud.go

    mkdir routes/$file_name
    echo -e "$header\n\npackage $file_name\n\n$initHandler" > routes/$file_name/handler.go
    echo -e "$header\n\npackage $file_name\n\n$initRoute" > routes/$file_name/route.go

    echo -e "$header\n\npackage model\n" > model/$file_name.go
}

function OpenModules() {
    zeditor --add routes/routes.go
    zeditor --add model/$file_name.go
    zeditor --add routes/$file_name/handler.go
    zeditor --add routes/$file_name/route.go
    zeditor --add crud/$file_name/crud.go
}

ExistFile=crud/$file_name/crud.go

if [ ! -f $ExistFile ];
then
    CreateLayout
    OpenModules
fi

if [ -f $ExistFile ];
then
    OpenModules
fi
