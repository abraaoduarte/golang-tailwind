package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/resume/internal/database"
	"example.com/resume/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := database.Connect()
    if err != nil {
        fmt.Println("Error trying to connect with database", err)
        return
    }

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("Error", err)
	}
    defer conn.Close(context.Background())

    // database.RunMigration()
    database.RunMigration()

	server := echo.New()
    server.Pre(middleware.AddTrailingSlash())
    server.Use(middleware.Logger())

    routes.RegisterRouters(conn, server)

    server.Static("/css", "/css")
    server.Static("/static", "static")

	err = server.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on http://localhost:%v", os.Getenv("PORT"))
}
