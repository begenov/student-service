package domain

type Courses struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Students []string `json:"students"`
}
