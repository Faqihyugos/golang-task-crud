package handlers

import (
	"strconv"

	"github.com/Faqihyugos/golang-task-crud/entities"
	"github.com/Faqihyugos/golang-task-crud/repositories"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskRepository repositories.TaskRepository
}

func NewTaskHandler(taskRepository repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepository: taskRepository}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	task := new(entities.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	if err := h.taskRepository.CreateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": task})
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.taskRepository.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tasks})
}

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	taskID, _ := strconv.Atoi(c.Params("id"))
	task, err := h.taskRepository.GetTaskByID(uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": task})
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	task := new(entities.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	taskId, _ := strconv.Atoi(c.Params("id"))
	task.ID = uint(taskId)

	if err := h.taskRepository.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": task})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	taskID, _ := strconv.Atoi(c.Params("id"))

	if err := h.taskRepository.DeleteTask(uint(taskID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Task successfully deleted"})
}
