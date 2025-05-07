package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Task struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("Hello again!")
	//fmt.Println("Test")
	app:=fiber.New()
	
	tasks:= []Task{}

	app.Get("/", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg":"Hi from Task Manager"})
	})

	app.Post("/api/tasks",func(c *fiber.Ctx) error{
		task := &Task{}

		if err := c.BodyParser(task)
		err!=nil{
			return err
		}

		if task.Body ==""{
			return c.Status(400).JSON(fiber.Map{
				"error":"Task must have a description"})
		}

		task.ID= len(tasks) +1
		tasks = append(tasks, *task)

		return c.Status(201).JSON(fiber.Map{
			"msg":"Task has been added succesfully!"})



	})


	log.Fatal(app.Listen(":4000"))
}