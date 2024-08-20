package routes

import (
	"os"
	"time"

	"github.com/backend/config"
	"github.com/backend/routes/favorites"
	"github.com/backend/routes/login"
	"github.com/backend/routes/users"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handlerHealth(c *fiber.Ctx) error {
	return c.JSON(
		map[string]string{
			"time":                time.Now().String(),
			"message":             os.Getenv("MESSAGE"),
			"service_name":        os.Getenv("NAME"),
			"service_version":     os.Getenv("VERSION"),
			"service_environment": os.Getenv("APP"),
		},
	)
}

func InitRoutes(f *fiber.App, db *gorm.DB) {
	app := f.Group("/api")
	health(f)

	favorites.InitFavoriteRoutes(app, db)
	login.InitLoginRoutes(app, db)
	users.InitUserRoutes(app, db)

	f.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JwtSecretKey())},
	}))

}

func health(f *fiber.App) {
	f.Get("/api/health", handlerHealth)
}
