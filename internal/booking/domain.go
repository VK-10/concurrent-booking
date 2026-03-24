package booking

import (
	"errors"
	"time"
)

var (
	ErrorSeatAlreadyBooked = errors.New("seat is already booked")
)

//Booking represents a confirmed seat reservation

type Booking struct {
	ID        string
	MovieID   string
	SeatID    string
	UserID    string
	Status    string
	ExpiredAt time.Time
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
