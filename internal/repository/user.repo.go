package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
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
