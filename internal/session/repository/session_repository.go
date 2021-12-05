package repository

import (
	"errors"
	"fmt"
	"github.com/TechnoHandOver/backend/internal/models"
)

type SessionRepository struct {
	db map[string]uint32
}

func NewSessionRepositoryImpl() *SessionRepository {
	return &SessionRepository{
		db: make(map[string]uint32),
	}
}

func (sessionRepository *SessionRepository) Insert(session *models.Session) (*models.Session, error) {
	if userId, ok := sessionRepository.db[session.Id]; ok {
		return nil, errors.New(fmt.Sprintf("Already exists: id = %d\n", userId))
	}

	sessionRepository.db[session.Id] = session.UserId
	return session, nil
}

func (sessionRepository *SessionRepository) Select(id string) (*models.Session, error) {
	userId, ok := sessionRepository.db[id]
	if !ok {
		return nil, errors.New("Doesn't exists\n")
	}

	return &models.Session{
		Id:     id,
		UserId: userId,
	}, nil
}
