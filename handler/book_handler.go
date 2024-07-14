package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"GoDeveloperTest/core"
	m "GoDeveloperTest/models"
)

type bookHandler struct {
	bookServ core.BookService
}

func NewBookHandler(bookServ core.BookService) *bookHandler {
	return &bookHandler{bookServ: bookServ}
}

func (h *bookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.bookServ.GetBooks()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(books)
}

func (h *bookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book, err := h.bookServ.GetBookById(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(book)
}

func (h *bookHandler) CreateBook(c *fiber.Ctx) error {
	var book m.Book

	err := c.BodyParser(&book)
	if err != nil {
		fmt.Print(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = h.bookServ.CreateBook(book)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"message": "Created book successful"})
}

func (h *bookHandler) UpdateBook(c *fiber.Ctx) error {
	var book m.Book

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
		
	}
	err = c.BodyParser(&book)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book.ID = uint(id)

	err = h.bookServ.UpdateBook(book)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"message": "Updated book successful"})
}

func (h *bookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = h.bookServ.DeleteBook(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"message": "Deleted book successful"})
}
