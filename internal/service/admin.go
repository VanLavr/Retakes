package service

import (
	"github.com/VanLavr/Retakes/internal/controller"
)

type AdminRepo interface{}

type adminService struct{}

func NewAdminService(r AdminRepo) controller.AdminService {
	return &adminService{}
}

func (a adminService) UploadRetakeData() {
	panic("not implemented")
}

func (a adminService) UploadStudentRetakes() {
	panic("not implemented")
}

func (a adminService) UploadTeacherRetakes() {
	panic("not implemented")
}
