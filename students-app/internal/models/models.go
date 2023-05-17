package models

type Student struct {
	ID      int      `json:"id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	GPA     float64  `json:"-"`
	Courses []string `json:"courses"`
}
