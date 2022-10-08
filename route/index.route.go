package route

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-gorm/handler"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerRead)
}
