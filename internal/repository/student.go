package repository

import (
	"github.com/VanLavr/Retakes/internal/service"
)

type studentRepo struct{}

func NewStudentRepo() service.StudentRepo {
	return &studentRepo{}
}

func (a studentRepo) UploadRetakeData() {
	panic("not implemented")
}

func (a studentRepo) UploadStudentRetakes() {
	panic("not implemented")
}

func (a studentRepo) UploadTeacherRetakes() {
	panic("not implemented")
}
