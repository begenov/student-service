package models

type Student struct {
	ID      int      `json:"id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	GPA     float64  `json:"gpa"`
	Courses []string `json:"courses"`
}

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
