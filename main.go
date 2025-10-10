package main

import (
    "log"
    "time"
    "math/rand"
    "github.com/gofiber/fiber/v3"
)

type Response struct {
	Message string "json:\"message\""
	Timestamp int64 "json:\"timestamp\""
	Favorite_Color string "json:\"favorite_color\""
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New()
    colors := []string{"black", "white", "red", "green", 
    		       "yellow", "blue", "brown", "orange",
		       "pink", "purple", "gray"}

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
	response := Response{
			Message:  "My name is Christopher Gemperle",
			Timestamp: time.Now().UnixMilli(),
			Favorite_Color: colors[rand.Intn(len(colors))],
	}
	return c.JSON(response)
    })

    // Start the server on port 80
    log.Fatal(app.Listen(":80"))
}
