package main

import (
    "log"
    "time"
    "fmt"
    "github.com/gofiber/fiber/v3"
    "github.com/goccy/go-json"
)

type Response struct {
	Message string "json:\"message\""
	Timestamp int64 "json:\"timestamp\""
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New(fiber.Config{
	    JSONEncoder: json.Marshal, // Minifies JSON String
	})

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
	response := Response{
			Message:  "My name is Christopher Gemperle",
			Timestamp: time.Now().UnixMilli(),
	}

	minifiedJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshaling:", err)
	}
	fmt.Println("Minified Output:", string(minifiedJSON))

	// Return the product as a JSON object
	return c.JSON(minifiedJSON)
    })

    // Start the server on port 80
    log.Fatal(app.Listen(":80"))
}
