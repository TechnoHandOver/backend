package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/notification"
	"time"
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
(SELECT user_.id, user_.vk_id, user_.name, user_.avatar
FROM user_
JOIN (SELECT route.user_author_id FROM route_tmp
    JOIN route ON route_tmp.id = route.id
WHERE route.user_author_id != $1 AND
      to_tsvector('russian', route.loc_dep) @@ plainto_tsquery('russian', $2) AND
      to_tsvector('russian', route.loc_arr) @@ plainto_tsquery('russian', $3) AND
      route.min_price <= $4 AND
      route_tmp.date_time_dep <= $5 AND
      route_tmp.date_time_arr >= $5) AS "route_tmp_"
    ON route_tmp_.user_author_id = user_.id)
UNION
(SELECT user_.id, user_.vk_id, user_.name, user_.avatar
FROM user_
JOIN (SELECT route.user_author_id FROM route_perm
    JOIN route ON route_perm.id = route.id
WHERE route.user_author_id != $1 AND
      to_tsvector('russian', route.loc_dep) @@ plainto_tsquery('russian', $2) AND
      to_tsvector('russian', route.loc_arr) @@ plainto_tsquery('russian', $3) AND
      route.min_price <= $4 AND
      route_perm.day_of_week = extract(ISODOW FROM $2) AND
      to_timestamp(to_char(route_perm.time_dep, 'HH24:mi'), 'HH24:mi') >= $5 AND
      to_timestamp(to_char(route_perm.time_arr, 'HH24:mi'), 'HH24:mi') <= $5) AS "route_perm_"
    ON route_perm_.user_author_id = user_.id)`

	rows, err := notificationRepository.db.Query(query, ad.UserAuthorId, ad.LocDep, ad.LocArr, ad.MinPrice,
		time.Time(ad.DateTimeArr))
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
