package usecase

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/session"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/google/uuid"
)

type SessionUsecase struct {
	sessionRepository session.Repository
}

func NewSessionUsecaseImpl(sessionRepository session.Repository) *SessionUsecase {
	return &SessionUsecase{
		sessionRepository: sessionRepository,
	}
}

func (sessionUsecase *SessionUsecase) Create(userVkId uint32) *response.Response {
	session_ := &models.Session{
		Id:       uuid.NewString(),
		UserVkId: userVkId,
	}

	session_, err := sessionUsecase.sessionRepository.Insert(session_)
	if err != nil {
		return response.NewEmptyResponse(consts.NotFound)
	}

	return response.NewResponse(consts.OK, session_)
}

func (sessionUsecase *SessionUsecase) Get(id string) *response.Response {
	session_, err := sessionUsecase.sessionRepository.Select(id)
	if err != nil {
		return response.NewEmptyResponse(consts.NotFound)
	}

	return response.NewResponse(consts.OK, session_)
}
