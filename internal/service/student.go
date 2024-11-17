package service

import (
	"github.com/VanLavr/Retakes/internal/controller"
)

type StudentRepo interface{}

type studentService struct{}

func NewStudentService(r StudentRepo) controller.StudentService {
	return &studentService{}
}

func (a studentService) UploadRetakeData() {
	panic("not implemented")
}

func (a studentService) UploadStudentRetakes() {
	panic("not implemented")
}

func (a studentService) UploadTeacherRetakes() {
	panic("not implemented")
}
