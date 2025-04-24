package service

import (
    "taskMangementService/model"
    "taskMangementService/repository"
)

type TaskService struct {
    Repo *repository.TaskRepository
}

func (s *TaskService) CreateTask(task *model.Task) error {
    return s.Repo.Create(task)
}

func (s *TaskService) GetTasks(status string, page, limit int) ([]model.Task, error) {
    offset := (page - 1) * limit
    return s.Repo.GetAll(status, limit, offset)
}

func (s *TaskService) GetTaskByID(id uint) (model.Task, error) {
    return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *model.Task) error {
    return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
    return s.Repo.Delete(id)
}
