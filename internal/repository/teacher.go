package repository

import (
	"github.com/VanLavr/Retakes/internal/service"
)

type teacherRepo struct{}

func NewteacherRepo() service.TeacherRepo {
	return &teacherRepo{}
}

func (a teacherRepo) UploadRetakeData() {
	panic("not implemented")
}

func (a teacherRepo) UploadteacherRetakes() {
	panic("not implemented")
}

func (a teacherRepo) UploadTeacherRetakes() {
	panic("not implemented")
}
