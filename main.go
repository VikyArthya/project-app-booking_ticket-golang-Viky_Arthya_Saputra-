package main

import (
	"log"
	"net/http"
	"ticket-booking/database"
	"ticket-booking/router"
	"ticket-booking/util"
)

func main() {
	// Inisialisasi logger
	util.InitLogger()

	// Inisialisasi database
	database.InitDB()

	// Membuat router dan menambahkan routes
	r := router.NewRouter()

	// Menjalankan server pada port 3000
	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
