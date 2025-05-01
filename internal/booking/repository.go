package booking

import (
	"fmt"

	"github.com/darkphotonKN/flow/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookingRepository struct {
	DB *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) *BookingRepository {
	return &BookingRepository{
		DB: db,
	}
}

func (r *BookingRepository) Create(userId uuid.UUID, req CreateBookingRequest) error {
	query := `INSERT INTO bookings(user_id, start_date, end_date, status) VALUES (:userId, :startDate, :endDate, :status)`

	// create fields required for DB insert
	fields := map[string]interface{}{
		"userId":    userId,
		"startDate": req.StartDate,
		"endDate":   req.EndDate,
		"status":    req.Status,
	}

	fmt.Printf("fields: %+v", fields)

	_, err := r.DB.NamedExec(query, fields)

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) GetById(userId uuid.UUID, id uuid.UUID) (*models.Booking, error) {
	// One to Many example
	query := `
	SELECT 
		bookings.id,
		bookings.start_date,
		bookings.end_date,
		bookings.status
	FROM bookings
	JOIN users ON users.id = bookings.user_id
	WHERE bookings.id = $1 AND bookings.user_id = $2
	`

	var booking models.Booking

	err := r.DB.Get(&booking, query, id, userId)

	if err != nil {
		return nil, err
	}

	return &booking, nil
}
