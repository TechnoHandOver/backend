package notification

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type Usecase interface {
	NotifySuitableUsers(ad *models.Ad) *response.Response
}
