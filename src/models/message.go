package models

type Message struct {
	Message_id          int
	Message_thread_code string
	Message             string
	Sender              string
	Timestamp           string
	Read_status         int
}
