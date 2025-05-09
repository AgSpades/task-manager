package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Task struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("Hello again!")
	app:=fiber.New()
	err:=godotenv.Load(".env")

	if err!=nil{
		log.Fatal("Error loading .env file!")
	}

	PORT :=os.Getenv("PORT")

	tasks:= []Task{}

	app.Get("/", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg":"Hi from Task Manager"})
	})

	//method for viewing list of tasks
	app.Get("/api/tasks",func(c*fiber.Ctx) error{
		return c.Status(200).JSON(tasks)
	})

	//method for adding tasks
	app.Post("/api/tasks",func(c *fiber.Ctx) error{
		task := &Task{}

		if err := c.BodyParser(task);err!=nil{
			return err
		}

		if task.Body ==""{
			return c.Status(400).JSON(fiber.Map{
				"error":"Task must have a description"})
		}

		task.ID= len(tasks) +1
		tasks = append(tasks, *task)

		return c.Status(201).JSON(task)

	})

	//method for marking tasks done
	app.Patch("api/tasks/:id",func(c*fiber.Ctx) error{
		id:=c.Params("id")

		for i, task:=range tasks{
			if fmt.Sprint(task.ID)==id{
				tasks[i].Completed= true
				return c.Status(200).JSON(tasks[i])
			}

		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	//method for deleting a task
	app.Delete("api/tasks/:id",func(c *fiber.Ctx) error{
		id:=c.Params("id")

		for i, task:=range tasks{
			if fmt.Sprint(task.ID)==id{
				tasks= append(tasks[:i], tasks[i+1:]... )
				
				return c.Status(200).JSON((fiber.Map{"success":"Deleted task succesfully"}))
			}

		}


		return c.Status(404).JSON((fiber.Map{"error":"Todo with valid id not found"}))
	})


	log.Fatal(app.Listen(":"+PORT))
}