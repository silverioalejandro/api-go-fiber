package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/silverioalejandro/api-go-fiber/database"
	"github.com/silverioalejandro/api-go-fiber/database/migration"
	"github.com/silverioalejandro/api-go-fiber/route"
)

func main() {
	//INITIAL DATABASE
	database.DatabaseInit()

	//RUN MIGRATION
	migration.RunMigrations()

	app := fiber.New()

	// INITIAL ROUTE

	route.RouteInit(app)

	app.Listen(":3000")
}
