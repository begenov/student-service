package models

type Course struct {
	ID         int    `json:"id"`
	Student_id []int  `json:"student_id"`
	Name       string `json:"name"`
}
