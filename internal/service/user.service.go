package service

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/kodacampmain/koda-b7-gin/internal/dto"
	"github.com/kodacampmain/koda-b7-gin/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
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

func (u *UserService) GetEmployees(ctx context.Context) ([]dto.Employees, error) {
	data, err := u.userRepo.FetchEmployeeList(ctx)
	if err != nil {
		return nil, err
	}
	var employees []dto.Employees
	for _, employee := range data {
		employees = append(employees, dto.Employees{
			Id:     employee.Id,
			Name:   employee.Name,
			Salary: employee.Salary,
		})
	}
	return employees, nil
}

func (u *UserService) NewEmployee(ctx context.Context, employee dto.NewEmployee) (dto.Employees, error) {
	result, err := u.userRepo.CreateNewEmployee(ctx, employee)
	if err != nil {
		return dto.Employees{}, err
	}
	newEmployee := dto.Employees{
		Id:   result.Id,
		Name: result.Name,
	}
	return newEmployee, nil
}

func (u *UserService) GetUserProfile(ctx context.Context, id int) (dto.User, error) {
	res, err := u.userRepo.FetchUserProfile(ctx, id)
	if err != nil {
		return dto.User{}, err
	}
	return dto.User{
		Id:        res.Id,
		Username:  res.Username,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
