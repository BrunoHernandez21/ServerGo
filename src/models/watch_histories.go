package models

type Watch_histories struct {
	Watch_history_id   int
	Course_id          int
	Student_id         int
	Watching_lesson_id int
	Date_added         string
	Date_updated       string
	Quiz_result        string
}
