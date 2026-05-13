package dto

type MoviesQuery struct {
	Title string   `form:"title" json:"title"`
	Genre []string `form:"genre" json:"genre"`
}
