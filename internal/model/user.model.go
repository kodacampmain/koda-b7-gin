package model

type Employee struct {
	Id     int    `db:"id"`
	Name   string `db:"employee_name"`
	DepId  int    `db:"department_id"`
	Salary int    `db:"salary"`
}
