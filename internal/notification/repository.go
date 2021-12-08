package notification

import "github.com/TechnoHandOver/backend/internal/models"

type Repository interface {
	SelectUsersByRoutesWithSuitableTimeInterval(ad *models.Ad) (*models.Users, error)
}
