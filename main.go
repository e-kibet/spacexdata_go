package main

import (
	"demo/database"
	"demo/book"
	"demo/database"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Print("Created the database connection")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)

	app.Get("/api/v1/books/:id", book.GetBook)

	app.Post("/api/v1/books", book.NewBooks)

	app.Delete("/api/v1/books/:id", book.DeleteBooks)
}

func main() {
	app := fiber.New()

	setUpRoutes(app)
	initDatabase()
	defer database.DBConn.Close()
	log.Fatal(app.Listen(":3200"))
}
