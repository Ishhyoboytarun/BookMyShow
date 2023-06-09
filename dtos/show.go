package dtos

import "time"

type CreateShowRequest struct {
	TimeSlot   string      `json:"time_slot"`
	TheatreID  int         `json:"theatre_id"`
	Screen     int         `json:"screen"`
	MovieID    int         `json:"movie_id"`
	Features   []string    `json:"features"`
	Seats      []int       `json:"seats"`
	BookedSeat map[int]int `json:"booked_seats"`
	StartTime  *time.Time  `json:"start_time"`
	EndTime    *time.Time  `json:"end_time"`
}

type ShowResponse struct {
	ID         int              `json:"id"`
	TimeSlot   string           `json:"time_slot"`
	Theatre    *TheatreResponse `json:"theatre"`
	Screen     int              `json:"screen"`
	Movie      *MovieResponse   `json:"movie"`
	Features   []string         `json:"features"`
	Seats      []*SeatResponse  `json:"seats"`
	BookedSeat map[int]int      `json:"booked_seats"`
}
