// main.go
package main

import (
    "fmt"
    "encoding/json"
    "net"
    "net/url"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func validateURL(s string) bool {
    if !(strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")) {
        return false
    }
    u, err := url.Parse(s)
    if err != nil {
        return false
    }

    fmt.Println(u.Host)

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

    app.Use(cors.New())

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
        c.Accepts("application/x-www-form-urlencoded")
        body := c.Body()
        fmt.Println(string(body))
        var data map[string]string
        json.Unmarshal(body, &data)
        fmt.Println(data)
        payloadUrl := strings.Replace(string(body), "url=", "", 1)
        if validateURL(payloadUrl) {
            doc := InsertURL(payloadUrl)
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
