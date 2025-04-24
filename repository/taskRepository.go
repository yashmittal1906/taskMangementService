package repository

import (
	"taskMangementService/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
    DB *gorm.DB
}

func (r *TaskRepository) Create(task *model.Task) error {
    return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAll(limit, offset int) ([]model.Task, error) {
    var tasks []model.Task
    query := r.DB.Limit(limit).Offset(offset)
    err := query.Find(&tasks).Error
    return tasks, err
}

func (r *TaskRepository) GetByID(id uint) (model.Task, error) {
    var task model.Task
    err := r.DB.First(&task, id).Error
    return task, err
}

func (r *TaskRepository) Update(task *model.Task) error {
    return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
    return r.DB.Delete(&model.Task{}, id).Error
}
