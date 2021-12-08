package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/notification"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepositoryImpl(db *sql.DB) notification.Repository {
	return &NotificationRepository{
		db: db,
	}
}

func (notificationRepository *NotificationRepository) SelectUsersByRoutesWithSuitableTimeInterval(ad *models.Ad) (*models.Users, error) {
	const query = `
SELECT DISTINCT ON (user_.id) user_.id, user_.vk_id, user_.name, user_.avatar
FROM user_
JOIN route_tmp
    ON route_tmp.date_time_dep < $1 AND route_tmp.date_time_arr >= $1
JOIN route_perm
    ON route_perm.day_of_week = extract(ISODOW FROM $1) AND route_perm.time_dep `

	rows, err := notificationRepository.db.Query(query, ad.DateTimeArr)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	users := make(models.Users, 0)
	for rows.Next() {
		user := new(models.User)
		if err := rows.Scan(&user.Id, &user.VkId, &user.Name, &user.Avatar); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}
