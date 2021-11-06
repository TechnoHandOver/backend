package session

import "github.com/TechnoHandOver/backend/internal/models"

type Repository interface {
	Insert(session *models.Session) (*models.Session, error)
	Select(id string) (*models.Session, error)
}
