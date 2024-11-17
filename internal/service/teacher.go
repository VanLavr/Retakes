package service

import (
	"github.com/VanLavr/Retakes/internal/controller"
)

type TeacherRepo interface{}

type teacherService struct{}

func NewTeacherService(r TeacherRepo) controller.TeacherService {
	return &teacherService{}
}

func (a teacherService) UploadRetakeData() {
	panic("not implemented")
}

func (a teacherService) UploadTeacherRetakes() {
	panic("not implemented")
}
