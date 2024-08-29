package main

import (
	// "ariga.io/atlas-provider-gorm/gormschema"
	"fmt"

	"github.com/Faqihyugos/golang-task-crud/handlers"
	"github.com/Faqihyugos/golang-task-crud/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// this migrate with gorm and atlas
	// stmts, err := gormschema.New("mysql").Load(&entities.Task{})
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
	// 	os.Exit(1)
	// }
	// io.WriteString(os.Stdout, stmts)

	app := fiber.New()
	dsn := "root:pass@tcp(localhost:3306)/db_crud?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if db != nil {
		fmt.Println("connected")
	}

	if err != nil {
		panic("failed to connect database")
	}

	taskRepository := repositories.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepository)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/tasks", taskHandler.CreateTask)
	app.Get("/tasks", taskHandler.GetAllTasks)
	app.Get("/tasks/:id", taskHandler.GetTaskByID)
	app.Put("/tasks/:id", taskHandler.UpdateTask)
	app.Delete("/tasks/:id", taskHandler.DeleteTask)

	app.Listen(":3001")
}
