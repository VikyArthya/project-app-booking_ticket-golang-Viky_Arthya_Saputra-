package middleware

import (
	"net/http"
	"ticket-booking/util"
	"time"

	"go.uber.org/zap"
)

// RequestLogger middleware untuk logging setiap request dan response
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mencatat waktu request diterima
		start := time.Now()

		// Menjalankan handler berikutnya
		next.ServeHTTP(w, r)

		// Mencatat log setelah request selesai
		util.Logger.Info("Request processed",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
