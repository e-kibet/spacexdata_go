package main

import (
	"demo/book"
	"demo/controllers"
	"log"
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	app.Get("/api/v1/books", book.GetBooks)

	app.Get("/api/v1/books/:id", book.GetBook)

	app.Post("/api/v1/books", book.NewBooks)

	app.Delete("/api/v1/books/:id", book.DeleteBooks)

	userRepo := controllers.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	infoRepo := controllers.NewInfo()
	r.GET("/infos", infoRepo.GetInfos)

}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the spacexdata")
	})
	setUpRoutes(app)
	app.Get("/swagger/*", swagger.Handler)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "http://example.com/doc.json",
		DeepLinking: true,
	}))
	log.Fatal(app.Listen(":3200"))
}
