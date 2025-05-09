package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Task struct{
	ID int `json:"_id" bson:"_id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

var collection *mongo.Collection

func main(){
	fmt.Println("Here we go again!")

	err:= godotenv.Load(".env")

	if err!= nil{
		log.Fatal("Error loading .env file",err)
	}

	MONGODB_URI:=os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err:= mongo.Connect(context.Background(),clientOptions)

	if err!=nil{
		log.Fatal(err)
	}

	err=client.Ping(context.Background(),nil)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	collection = client.Database("tmanager_db").Collection("tasks")

	app:= fiber.New()

	// endpoints for the API

	app.Get("/api/tasks", getTasks)
	app.Post("/api/tasks", createTask)
	app.Patch("/api/tasks/:id", updateTask)
	app.Delete("/api/tasks/:id", deleteTask)


	PORT := os.Getenv("PORT")

	if PORT==""{
		PORT="5000"
	}
	log.Fatal(app.Listen(":"+PORT))


}

func getTasks(c *fiber.Ctx) error {}
func createTask(c *fiber.Ctx) error {}
func updateTask(c *fiber.Ctx) error {}
func deleteTask(c *fiber.Ctx) error {}