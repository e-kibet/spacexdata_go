package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Return all books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Return One book")
}

func NewBooks(c *fiber.Ctx) error {
	return c.SendString("add new book")
}

func DeleteBooks(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
