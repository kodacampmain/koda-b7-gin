package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b7-gin/internal/dto"
	"github.com/kodacampmain/koda-b7-gin/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FetchEmployeeList(ctx context.Context) ([]model.Employee, error) {
	// jalankan query
	sql := "SELECT id, employee_name, department_id, salary FROM employees LIMIT 100"
	rows, err := u.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Employee
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.DepId, &employee.Salary); err != nil {
			return nil, err
		}
		result = append(result, employee)
	}
	return result, nil

}

func (u *UserRepository) CreateNewEmployee(ctx context.Context, employee dto.NewEmployee) (model.Employee, error) {
	sql := "INSERT INTO employees (employee_name, department_id, salary) VALUES ($1, $2, $3) RETURNING id, employee_name"
	// args := []any{1,2,3,4,5,6}
	// $1 => parameterized query
	row := u.db.QueryRow(ctx, sql, employee.Name, employee.DepId, employee.Salary)
	var result model.Employee
	if err := row.Scan(&result.Id, &result.Name); err != nil {
		return model.Employee{}, err
	}
	return result, nil
}

func (u *UserRepository) FetchUserProfile(ctx context.Context, id int) (model.User, error) {
	sql := "SELECT id, username, created_at, updated_at FROM users WHERE id = $1"
	args := []any{id}
	var user model.User
	if err := u.db.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return model.User{}, err
	}
	return user, nil
}
