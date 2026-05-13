package service

import (
	"log"

	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

type UserServiceMock struct{}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (u *UserServiceMock) PrintUser(body dto.UsersBody) {
	log.Println("data body masuk")
}
