package repository

import (
	"log"
	"ticket-booking/database"
	"ticket-booking/model"
)

func CreateBooking(booking model.Booking) (int, error) {
	var bookingID int
	query := `INSERT INTO bookings (name, email, confirm_email, phone, date, number_of_ticket, message)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := database.DB.QueryRow(query, booking.Name, booking.Email, booking.ConfirmEmail, booking.Phone, booking.Date, booking.NumberOfTicket, booking.Message).Scan(&bookingID)
	if err != nil {
		log.Println("Error creating booking: ", err)
		return 0, err
	}
	return bookingID, nil
}

func GetBookingByID(id int) (model.Booking, error) {
	var booking model.Booking
	query := `SELECT id, name, email, confirm_email, phone, date, number_of_ticket, message FROM bookings WHERE id=$1`
	err := database.DB.QueryRow(query, id).Scan(&booking.ID, &booking.Name, &booking.Email, &booking.ConfirmEmail, &booking.Phone, &booking.Date, &booking.NumberOfTicket, &booking.Message)
	if err != nil {
		return booking, err
	}
	return booking, nil
}

func GetAllBookings() ([]model.Booking, error) {
	rows, err := database.DB.Query(`SELECT id, name, email, confirm_email, phone, date, number_of_ticket, message FROM bookings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []model.Booking
	for rows.Next() {
		var booking model.Booking
		if err := rows.Scan(&booking.ID, &booking.Name, &booking.Email, &booking.ConfirmEmail, &booking.Phone, &booking.Date, &booking.NumberOfTicket, &booking.Message); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func UpdateBooking(id int, booking model.Booking) error {
	query := `UPDATE bookings SET name=$1, email=$2, confirm_email=$3, phone=$4, date=$5, number_of_ticket=$6, message=$7 WHERE id=$8`
	_, err := database.DB.Exec(query, booking.Name, booking.Email, booking.ConfirmEmail, booking.Phone, booking.Date, booking.NumberOfTicket, booking.Message, id)
	return err
}

func DeleteBooking(id int) error {
	query := `DELETE FROM bookings WHERE id=$1`
	_, err := database.DB.Exec(query, id)
	return err
}
