package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Task struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
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

	defer client.Disconnect(context.Background())

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

func getTasks(c *fiber.Ctx) error {
	var tasks []Task
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}

	defer cursor.Close(context.Background())

	for(cursor.Next(context.Background())){
		var task Task
		if err := cursor.Decode(&task); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to decode task",
			})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func createTask(c *fiber.Ctx) error {
	task:=new(Task)

	if err:= c.BodyParser(task); err!=nil{
		log.Fatal(err)
	}

	if task.Body == ""{
		return c.Status(400).JSON(fiber.Map{
			"error": "Task body is required",
		})
	}

	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	task.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(task)
}
func updateTask(c *fiber.Ctx) error {
	id:=c.Params("id")
	objectID,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}


	filter :=bson.M{"_id":objectID} // filter to find the task
	// update the task with the given ID
	update:=bson.M{"$set":bson.M{"completed":true}}

	_,err=collection.UpdateOne(context.Background(),filter,update)

	if err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
func deleteTask(c *fiber.Ctx) error {
	id:=c.Params("id")
	objectID,err:=primitive.ObjectIDFromHex(id)

	if err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid task ID",
		})
	}

	filter :=bson.M{"_id":objectID} // filter to find the task
	// delete the task with the given ID
	_,err=collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}