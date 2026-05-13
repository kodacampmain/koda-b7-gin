package dto

import "time"

type UsersBody struct {
	// key datatype `tag`
	Fullname string    `json:"nama_lengkap,omitempty" form:"nl"`
	Email    string    `json:"surel" binding:"required,email"`
	Age      int       `json:"umur"`
	Dob      time.Time `json:"ttl"`
}

type UsersUri struct {
	Id   int    `uri:"id" json:"id"`
	Slug string `uri:"slug" json:"slug"`
}

type UsersHeader struct {
	ContentType string `header:"Content-Type" binding:"required"`
}

type Employees struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}
