package main

import (
    "database/sql"
    "os"
    "log"

    "github.com/gofiber/fiber/v2"
    _ "github.com/libsql/libsql-client-go/libsql"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    TURSO_TOKEN := os.Getenv("TURSO_TOKEN")
    TURSO_DB := os.Getenv("TURSO_DB")
    var dbUrl = "libsql://" + TURSO_DB + ".turso.io?authToken=" + TURSO_TOKEN
    log.Println("Connecting to " + dbUrl)
    db, err := sql.Open("libsql", dbUrl)
    if err != nil {
      log.Fatal(os.Stderr, "Error opening database: %v\n", err)
      os.Exit(1)
    }

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        var version string
        err = db.QueryRow("SELECT sqlite_version()").Scan(&version)
        if err != nil {
            log.Fatal(err)
        }
        return c.SendString("Hello, World! " + version)
    })

    log.Fatal(app.Listen(":3000"))
}
