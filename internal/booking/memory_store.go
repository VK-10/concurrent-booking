package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}
func (s *MemoryStore) Book(b Booking) error {
	// seat is taken
	if _, exists := s.bookings[b.SeatID]; exists == true {
		return ErrorSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil

}
func (s *MemoryStore) ListBookings(movieID string) []Booking {
	// filtering bookings by movieID(to check the seats at the time of boooking the that movie)
	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
