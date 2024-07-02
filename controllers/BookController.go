package controllers

import (
	"awesomeProject/data/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
)

var Books = []*models.Book{
	{1, "The Lord of Rings"},
	{2, "Harry Potter"},
}

func GetAllBooks(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"books":   Books,
	})
}

func PostBook(c *fiber.Ctx) error {
	/*
		var body fiber.Request
		err := c.BodyParser(&body)
		if err != nil {
			return err
		}
		fmt.Println(body)*/
	var body = new(models.Book)

	if err := c.BodyParser(body); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error()) //this returns nil, avoid
		return err
	}
	log.Debug(body.Id)
	log.Debug(body.Title)
	Books = append(Books, body)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"books":   Books,
	})

}

func DeleteBookById(ctx *fiber.Ctx) error {
	// Extraction
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID",
		})
	}

	index := -1
	for i, book := range Books {
		if book.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	Books = append(Books[:index], Books[index+1:]...)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
