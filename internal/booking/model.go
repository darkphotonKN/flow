package booking

import "time"

type CreateBookingRequest struct {
	Status    string    `json:"status"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}
