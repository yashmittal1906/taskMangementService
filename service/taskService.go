package service

import (
	"taskMangementService/model"
	"taskMangementService/repository"
)

type TaskService struct {
    Repo *repository.TaskRepository
	userClient *UserClient 
}

func (s *TaskService) CreateTask(task *model.Task) error {
	// Authenticate and get user details

	// user, err := (*s.userClient).ValidateAndGetUserByID(task.UserId) err != nil {
    // 	return err 
	// }	

    if err := s.Repo.Create(task); err != nil {
    	return err 
	}	

	// Notify users through emal and sms using kafka
	// (*s.userClient).notifyUser(user, task)
	return nil
}

func (s *TaskService) GetTasks(status string, page, limit int) ([]model.Task, error) {
	// Authenticate and get user details

	// user, err := (*s.userClient).ValidateAndGetUserByID(task.UserId) err != nil {
    // 	return err 
	// }

    offset := (page - 1) * limit
    return s.Repo.GetAll(status, limit, offset)
}

func (s *TaskService) GetTaskByID(id uint) (model.Task, error) {
	// Authenticate and get user details

	// user, err := (*s.userClient).ValidateAndGetUserByID(task.UserId) err != nil {
    // 	return err 
	// }

    return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *model.Task) error {
	// Authenticate and get user details

	// user, err := (*s.userClient).ValidateAndGetUserByID(task.UserId) err != nil {
    // 	return err 
	// }

    return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	// Authenticate and get user details

	// user, err := (*s.userClient).ValidateAndGetUserByID(task.UserId) err != nil {
    // 	return err 
	// }

    return s.Repo.Delete(id)
}
