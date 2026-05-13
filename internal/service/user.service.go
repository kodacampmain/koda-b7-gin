package service

import (
	"errors"
	"log"
	"regexp"

	"github.com/kodacampmain/koda-b7-gin/internal/dto"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) PrintUser(body dto.UsersBody) {
	// proses logika i.e. validasi
	log.Printf("\nNama: %s\nEmail: %s\nAge: %d\nDob: %s\n", body.Fullname, body.Email, body.Age, body.Dob.String())
}

func (u *UserService) ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if valid := regex.MatchString(email); !valid {
		return errors.New("invalid email format")
	}
	return nil
}
