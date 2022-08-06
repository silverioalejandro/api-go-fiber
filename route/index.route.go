package route 

import "github.com/gofiber/fiber/v2"

func RouteInit(r *fiber.App) {
	r.Get(path:"/user", handler.UserHandlerRead)
}