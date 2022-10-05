package main

import (
	"log"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
	mongo.NewClient
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) {

	})
	app.Post("/employee", func(c *fiber.Ctx) {

	})
	app.Put("/employee/:id", func(c *fiber.Ctx) {

	})
	app.Delete("/employee/:id", func(c *fiber.Ctx) {

	})
}
