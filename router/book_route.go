package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"

	"GoDeveloperTest/core"
	"GoDeveloperTest/handler"
	"GoDeveloperTest/repository"
)

func NewRouter() {
	app := fiber.New()

	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	bookRepository := repository.NewBookDB(db)
	bookService := core.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	app.Get("/api/books", bookHandler.GetBooks)
	app.Get("/api/books/:id", bookHandler.GetBookById)
	app.Post("/api/books", bookHandler.CreateBook)
	app.Put("/api/books/:id", bookHandler.UpdateBook)
	app.Delete("/api/books/:id", bookHandler.DeleteBook)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
