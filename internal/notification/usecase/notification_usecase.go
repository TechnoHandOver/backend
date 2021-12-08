package usecase

import (
	"fmt"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/notification"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"log"
	"net/http"
	"time"
)

type NotificationUsecase struct {
	notificationRepository notification.Repository
}

func NewNotificationUsecaseImpl(notificationRepository notification.Repository) notification.Usecase {
	return &NotificationUsecase{
		notificationRepository: notificationRepository,
	}
}

func (notificationUsecase *NotificationUsecase) NotifySuitableUsers(ad *models.Ad) *response.Response {
	users, err := notificationUsecase.notificationRepository.SelectUsersByRoutesWithSuitableTimeInterval(ad)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	client := &http.Client{
		Transport: &http.Transport{ //TODO: настроить
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	var anyErrorLogged = false
	for _, user := range *users {
		response_, err := client.Get(fmt.Sprintf("https://handover.space/bot/schedule?user_id=%d", user.Id))
		if err != nil {
			return response.NewErrorResponse(consts.InternalError, err)
		}
		if response_.StatusCode != http.StatusOK && !anyErrorLogged {
			log.Println("lobaevni: ", "cannot access vk bot: response code = ", response_.StatusCode)
			anyErrorLogged = true
		}
	}

	return response.NewEmptyResponse(consts.OK)
}
