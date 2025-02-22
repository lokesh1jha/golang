package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lokesh1jha/go-fiber-crm-basic/database"
	"github.com/lokesh1jha/go-fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
	app.Put("/api/v1/lead/:id", lead.UpdateLead)
}

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
	database.DB.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Use(logger.New())
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DB.Close()
}
