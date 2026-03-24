package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}
func (s *ConcurrentStore) Book(b Booking) error {

	s.Lock()
	defer s.Unlock()

	// seat is taken
	if _, exists := s.bookings[b.SeatID]; exists == true {
		return ErrorSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil

}
func (s *ConcurrentStore) ListBookings(movieID string) []Booking {

	s.RLock()
	defer s.RUnlock()

	// filtering bookings by movieID(to check the seats at the time of boooking the that movie)
	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
