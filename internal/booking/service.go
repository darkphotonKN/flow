package booking

import (
	"github.com/darkphotonKN/flow/internal/models"
	"github.com/google/uuid"
)

type BookingService struct {
	Repo *BookingRepository
}

func NewBookingService(repo *BookingRepository) *BookingService {
	return &BookingService{
		Repo: repo,
	}
}

func (s *BookingService) GetBookingByIdService(userId uuid.UUID, id uuid.UUID) (*models.Booking, error) {
	return s.Repo.GetById(userId, id)
}

func (s *BookingService) CreateBookingService(userId uuid.UUID, req CreateBookingRequest) error {
	return s.Repo.Create(userId, req)
}
