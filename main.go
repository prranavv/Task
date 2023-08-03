package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/prranavv/task/database"
)

func initapp() error {
	err := loadEnv()
	if err != nil {
		return err
	}
	err = database.StartMongodb()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initapp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongodb()
	app := generateapp()
	//get the port from the env
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func loadEnv() error {
	goenv := os.Getenv("GO_ENV")
	if goenv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateapp() *fiber.App {
	app := fiber.New()
	//create health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	return app
}
