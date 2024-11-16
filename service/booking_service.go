package service

import (
	"ticket-booking/model"
	"ticket-booking/repository"
)

func CreateBookingService(booking model.Booking) (int, error) {
	return repository.CreateBooking(booking)
}

func GetBookingServiceByID(id int) (model.Booking, error) {
	return repository.GetBookingByID(id)
}

func GetAllBookingsService() ([]model.Booking, error) {
	return repository.GetAllBookings()
}

func UpdateBookingService(id int, booking model.Booking) error {
	return repository.UpdateBooking(id, booking)
}

func DeleteBookingService(id int) error {
	return repository.DeleteBooking(id)
}
