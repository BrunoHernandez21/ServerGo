package models

type Lesson struct {
	Id                                int
	Title                             string
	Duration                          string
	Course_id                         int
	Section_id                        int
	Video_type                        string
	Video_url                         string
	Date_added                        int
	Last_modified                     int
	Lesson_type                       string
	Attachment                        string
	Attachment_type                   string
	Summary                           string
	Order                             int
	Video_type_for_mobile_application string
	Video_url_for_mobile_application  string
	Duration_for_mobile_application   string
	Is_free                           int
}
