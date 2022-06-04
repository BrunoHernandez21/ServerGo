package models

type Question struct {
	Id                int
	Quiz_id           int
	Title             string
	Type              string
	Number_of_options int
	Options           string
	Correct_answers   string
	Order             int
}
