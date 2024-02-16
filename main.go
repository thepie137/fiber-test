package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "good vibes", Author: "James"})
	books = append(books, Book{ID: 2, Title: "the alchemist", Author: "Boy"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)

	app.Listen(":8005")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")

}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Hello, World!",
		"description": "Hello, Description!",
	})
}
