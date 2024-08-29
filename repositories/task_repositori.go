package repositories

import (
	"github.com/Faqihyugos/golang-task-crud/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *entities.Task) error
	GetTaskByID(id uint) (*entities.Task, error)
	GetAllTasks() ([]entities.Task, error)
	UpdateTask(task *entities.Task) error
	DeleteTask(id uint) error
}

type taskRepositoryGorm struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryGorm{db: db}
}

func (r *taskRepositoryGorm) CreateTask(task *entities.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepositoryGorm) GetTaskByID(id uint) (*entities.Task, error) {
	var task entities.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepositoryGorm) GetAllTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepositoryGorm) UpdateTask(task *entities.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepositoryGorm) DeleteTask(id uint) error {
	return r.db.Delete(&entities.Task{}, id).Error
}
