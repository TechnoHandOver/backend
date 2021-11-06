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
	if userVkId, ok := sessionRepository.db[session.Id]; ok {
		return nil, errors.New(fmt.Sprintf("Already exists: id = %d\n", userVkId))
	}

	sessionRepository.db[session.Id] = session.UserVkId
	return session, nil
}

func (sessionRepository *SessionRepository) Select(id string) (*models.Session, error) {
	userVkId, ok := sessionRepository.db[id]
	if !ok {
		return nil, errors.New("Doesn't exists\n")
	}

	return &models.Session{
		Id:       id,
		UserVkId: userVkId,
	}, nil
}
