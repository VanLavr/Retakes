package repository

import (
	"github.com/VanLavr/Retakes/internal/service"
)

type adminRepo struct{}

func NewAdminRepo() service.AdminRepo {
	return &adminRepo{}
}

func (a adminRepo) UploadRetakeData() {
	panic("not implemented")
}

func (a adminRepo) UploadStudentRetakes() {
	panic("not implemented")
}

func (a adminRepo) UploadTeacherRetakes() {
	panic("not implemented")
}
