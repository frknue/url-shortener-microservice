package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Static("/", "./public")
  
  app.Post("/api/shorturl", func(c *fiber.Ctx) error {
    c.Accepts("application/json")
    body := c.Body()
    return c.SendString(string(body))
  })
  app.Listen(":3000")
  fmt.Println("Hello, World!")
}
