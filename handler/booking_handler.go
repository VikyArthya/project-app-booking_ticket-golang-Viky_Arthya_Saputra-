package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ticket-booking/model"
	"ticket-booking/service"
	"ticket-booking/util"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// CreateBookingHandler menangani permintaan untuk membuat booking baru
func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		util.Logger.Error("Invalid input", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	bookingID, err := service.CreateBookingService(booking)
	if err != nil {
		util.Logger.Error("Error creating booking", zap.Error(err))
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	// Logging sukses ketika booking berhasil dibuat
	util.Logger.Info("Booking created successfully",
		zap.Int("booking_id", bookingID),
	)

	response := map[string]interface{}{"id": bookingID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetBookingByIDHandler menangani permintaan untuk mendapatkan booking berdasarkan ID
func GetBookingByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Logger.Error("Invalid ID",
			zap.String("id", idStr),
			zap.Error(err),
		)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	booking, err := service.GetBookingServiceByID(id)
	if err != nil {
		util.Logger.Error("Booking not found",
			zap.Int("id", id),
			zap.Error(err),
		)
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Logging ketika booking ditemukan
	util.Logger.Info("Booking retrieved successfully",
		zap.Int("id", id),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

// GetAllBookingsHandler menangani permintaan untuk mendapatkan semua booking
func GetAllBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := service.GetAllBookingsService()
	if err != nil {
		util.Logger.Error("Error fetching bookings", zap.Error(err))
		http.Error(w, "Error fetching bookings", http.StatusInternalServerError)
		return
	}

	// Logging ketika semua booking berhasil diambil
	util.Logger.Info("Fetched all bookings successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// UpdateBookingHandler menangani permintaan untuk memperbarui booking berdasarkan ID
func UpdateBookingHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Logger.Error("Invalid ID",
			zap.String("id", idStr),
			zap.Error(err),
		)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var booking model.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		util.Logger.Error("Invalid input", zap.Error(err))
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = service.UpdateBookingService(id, booking)
	if err != nil {
		util.Logger.Error("Error updating booking",
			zap.Int("id", id),
			zap.Error(err),
		)
		http.Error(w, "Error updating booking", http.StatusInternalServerError)
		return
	}

	// Logging ketika booking berhasil diperbarui
	util.Logger.Info("Booking updated successfully",
		zap.Int("id", id),
	)

	w.WriteHeader(http.StatusOK)
}

// DeleteBookingHandler menangani permintaan untuk menghapus booking berdasarkan ID
func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Logger.Error("Invalid ID",
			zap.String("id", idStr),
			zap.Error(err),
		)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteBookingService(id)
	if err != nil {
		util.Logger.Error("Error deleting booking",
			zap.Int("id", id),
			zap.Error(err),
		)
		http.Error(w, "Error deleting booking", http.StatusInternalServerError)
		return
	}

	// Logging ketika booking berhasil dihapus
	util.Logger.Info("Booking deleted successfully",
		zap.Int("id", id),
	)

	w.WriteHeader(http.StatusNoContent)
}
