package domain

type Student struct {
	ID      int      `json:"id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	GPA     float64  `json:"gpa"`
	Courses []string `json:"courses"`
}

type UpdateStudentInput struct {
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	GPA     float64  `json:"gpa"`
	Courses []string `json:"courses"`
}

type CreateStudentInput struct {
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	GPA     float64  `json:"gpa"`
	Courses []string `json:"courses"`
}
