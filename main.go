// main.go
package main

import (
    "fmt"
    "encoding/json"
    "net"
    "net/url"
    "strings"

    "github.com/gofiber/fiber/v2"
)

func validateURL(s string) bool {
    if !(strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")) {
        return false
    }
    u, err := url.Parse(s)
    if err != nil {
        return false
    }

    ips, err := net.LookupIP(u.Host)
    if err != nil {
        return false
    }
    if len(ips) == 0 {
        return false
    }

    return true
}

func main() {
    app := fiber.New()

    app.Static("/", "./public")

    app.Get("api/shorturl/:shorturl", func(c *fiber.Ctx) error {
        shorturl := c.Params("shorturl")
        result := GetShortURL(shorturl)
        if result == nil {
            return c.SendStatus(404)
        }
        fmt.Println(result)
        return c.Redirect(result["original_url"])
    })

    app.Post("/api/shorturl", func(c *fiber.Ctx) error {
        c.Accepts("application/json")
        body := c.Body()
        var data map[string]string
        json.Unmarshal(body, &data)
        if validateURL(data["url"]) {
            doc := InsertURL(data["url"])
            return c.JSON(doc)
        } else {
            return c.JSON(fiber.Map{
                "error": "invalid URL",
            })
        }
    })

    fmt.Println("Listening on port 3000")
    app.Listen(":3000")
}
