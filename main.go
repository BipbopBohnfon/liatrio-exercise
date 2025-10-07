package main

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v3"
)

type Response struct {
	Message string "json:\"message\""
	Timestamp int64 "json:\"timestamp:\""
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
	response := Response{
			Message:  "My name is Christopher Gemperle",
			Timestamp: time.Now().Unix(),
	}
	// Return the product as a JSON object
	return c.JSON(response)
    })

    // Start the server on port 80
    log.Fatal(app.Listen(":80"))
}
