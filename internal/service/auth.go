package service

import (
	"github.com/VanLavr/Retakes/internal/controller"
)

type AuthRepo interface{}

type authService struct{}

func NewAuthService(r AuthRepo) controller.AuthService {
	return &authService{}
}

func (a authService) UploadRetakeData() {
	panic("not implemented")
}

func (a authService) UploadStudentRetakes() {
	panic("not implemented")
}

func (a authService) UploadTeacherRetakes() {
	panic("not implemented")
}
