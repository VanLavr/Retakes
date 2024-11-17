package repository

import (
	"github.com/VanLavr/Retakes/internal/service"
)

type authRepo struct{}

func NewAuthRepo() service.AuthRepo {
	return &authRepo{}
}

func (a authRepo) UploadRetakeData() {
	panic("not implemented")
}

func (a authRepo) UploadStudentRetakes() {
	panic("not implemented")
}

func (a authRepo) UploadTeacherRetakes() {
	panic("not implemented")
}
