package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	_, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	context.WithTimeout(context.Background(), 30*time.Second)

	return err
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
