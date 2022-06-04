package models

type Live_class struct {
	Id        int
	Course_id int
	Date      int
	Time      int

	Zoom_meeting_id       string
	Zoom_meeting_password string
	Note_to_students      string
}
