package model

type Booking struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ConfirmEmail   string `json:"confirm_email"`
	Phone          string `json:"phone"`
	Date           string `json:"date"`
	NumberOfTicket int    `json:"number_of_ticket"`
	Message        string `json:"message"`
}
