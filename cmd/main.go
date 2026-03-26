package main

import (
	"log"
	"net/http"

	"github.com/VK-10/concurrent-seat-booking/internal/booking"
	"github.com/VK-10/concurrent-seat-booking/internal/utils"
	"github.com/redis/go-redis/v9"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	}))
	svc := booking.NewService(store)

	bookingHandler := booking.Newhandler(svc)

	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /movies/{movieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)

	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

var movies = []movieResponse{
	{ID: "dhurandhar", Title: "Dhurandhar", Rows: 5, SeatsPerRow: 8},
	{ID: "project-hail-mary", Title: "Project Hail Mary", Rows: 5, SeatsPerRow: 5},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, movies)
}

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}
