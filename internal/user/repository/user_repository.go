package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/user"
	"strconv"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) Insert(user_ *models.User) (*models.User, error) {
	const query = "INSERT INTO user_ (vk_id, name, avatar) VALUES ($1, $2, $3) RETURNING id, vk_id, name, avatar"

	if err := userRepository.db.QueryRow(query, user_.VkId, user_.Name, user_.Avatar).Scan(&user_.Id, &user_.VkId,
		&user_.Name, &user_.Avatar); err != nil {
		return nil, err
	}

	return user_, nil
}

func (userRepository *UserRepository) Select(id uint32) (*models.User, error) {
	const query = "SELECT id, vk_id, name, avatar FROM user_ WHERE id = $1"

	user_ := new(models.User)
	if err := userRepository.db.QueryRow(query, id).Scan(&user_.Id, &user_.VkId, &user_.Name,
		&user_.Avatar); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return user_, nil
}

func (userRepository *UserRepository) SelectByVkId(vkId uint32) (*models.User, error) {
	const query = "SELECT id, vk_id, name, avatar FROM user_ WHERE vk_id = $1"

	user_ := new(models.User)
	var avatar sql.NullString
	if err := userRepository.db.QueryRow(query, vkId).Scan(&user_.Id, &user_.VkId, &user_.Name, &avatar); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}
	if avatar.Valid {
		user_.Avatar = avatar.String
	}

	return user_, nil
}

func (userRepository *UserRepository) Update(user_ *models.User) (*models.User, error) {
	const queryStart = "UPDATE user_ SET "
	const queryName = "name"
	const queryAvatar = "avatar"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = "WHERE id = $1 RETURNING id, vk_id, name, avatar"

	query := queryStart
	queryArgs := make([]interface{}, 0)
	queryArgs = append(queryArgs, user_.Id)

	if user_.Name != "" {
		query += queryName + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, user_.Name)
	}

	if user_.Avatar != "" {
		query += queryAvatar + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, user_.Avatar)
	}

	if len(queryArgs) == 1 {
		return userRepository.Select(user_.Id)
	}

	query = query[:len(query)-2] + queryEnd

	updatedUser := new(models.User)
	if err := userRepository.db.QueryRow(query, queryArgs...).Scan(&updatedUser.Id, &updatedUser.VkId,
		&updatedUser.Name, &updatedUser.Avatar); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return updatedUser, nil
}

func (userRepository *UserRepository) InsertRouteTmp(routeTmp *models.RouteTmp) (*models.RouteTmp, error) {
	const query = `
INSERT INTO view_route_tmp (user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr`

	if err := userRepository.db.QueryRow(query, routeTmp.UserAuthorVkId, routeTmp.LocDep, routeTmp.LocArr,
		routeTmp.MinPrice, time.Time(routeTmp.DateTimeDep), time.Time(routeTmp.DateTimeArr)).Scan(&routeTmp.Id,
		&routeTmp.UserAuthorVkId, &routeTmp.LocDep, &routeTmp.LocArr, &routeTmp.MinPrice, &routeTmp.DateTimeDep,
		&routeTmp.DateTimeArr); err != nil {
		return nil, err
	}

	return routeTmp, nil
}

func (userRepository *UserRepository) SelectRouteTmp(routeTmpId uint32) (*models.RouteTmp, error) {
	const query = `
SELECT id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp
WHERE id = $1`

	routeTmp := new(models.RouteTmp)
	if err := userRepository.db.QueryRow(query, routeTmpId).Scan(&routeTmp.Id, &routeTmp.UserAuthorVkId,
		&routeTmp.LocDep, &routeTmp.LocArr, &routeTmp.MinPrice, &routeTmp.DateTimeDep,
		&routeTmp.DateTimeArr); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return routeTmp, nil
}

func (userRepository *UserRepository) SelectRouteTmpArray() (*models.RoutesTmp, error) {
	const query = `
SELECT id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp
ORDER BY date_time_dep, date_time_arr, min_price DESC, id`

	rows, err := userRepository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	routesTmp := make(models.RoutesTmp, 0)
	for rows.Next() {
		routeTmp := new(models.RouteTmp)
		if err := rows.Scan(&routeTmp.Id, &routeTmp.UserAuthorVkId, &routeTmp.LocDep, &routeTmp.LocArr,
			&routeTmp.MinPrice, &routeTmp.DateTimeDep, &routeTmp.DateTimeArr); err != nil {
			return nil, err
		}

		routesTmp = append(routesTmp, routeTmp)
	}

	return &routesTmp, nil
}

func (userRepository *UserRepository) UpdateRouteTmp(routeTmp *models.RouteTmp) (*models.RouteTmp, error) {
	const queryStart = "UPDATE view_route_tmp SET "
	const queryLocDep = "loc_dep"
	const queryLocArr = "loc_arr"
	const queryMinPrice = "min_price"
	const queryDateTimeDep = "date_time_dep"
	const queryDateTimeArr = "date_time_arr"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = " WHERE id = $1 RETURNING id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr"

	query := queryStart
	queryArgs := make([]interface{}, 0)
	queryArgs = append(queryArgs, routeTmp.Id)

	if routeTmp.LocDep != "" {
		query += queryLocDep + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, routeTmp.LocDep)
	}

	if routeTmp.LocArr != "" {
		query += queryLocArr + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, routeTmp.LocArr)
	}

	if routeTmp.MinPrice != 0 {
		query += queryMinPrice + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, routeTmp.MinPrice)
	}

	if dateTimeDep := time.Time(routeTmp.DateTimeDep); !dateTimeDep.IsZero() {
		query += queryDateTimeDep + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, dateTimeDep)
	}

	if dateTimeArr := time.Time(routeTmp.DateTimeArr); !dateTimeArr.IsZero() {
		query += queryDateTimeArr + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, dateTimeArr)
	}

	if len(queryArgs) == 1 {
		return nil, consts.RepErrNothingToUpdate
	}

	query = query[:len(query)-2] + queryEnd

	if err := userRepository.db.QueryRow(query, queryArgs...).Scan(&routeTmp.Id, &routeTmp.UserAuthorVkId,
		&routeTmp.LocDep, &routeTmp.LocArr, &routeTmp.MinPrice, &routeTmp.DateTimeDep,
		&routeTmp.DateTimeArr); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return routeTmp, nil
}

func (userRepository *UserRepository) DeleteRouteTmp(routeTmpId uint32) (*models.RouteTmp, error) {
	const query = `
DELETE FROM view_route_tmp
WHERE id = $1
RETURNING id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr`

	routeTmp := new(models.RouteTmp)
	if err := userRepository.db.QueryRow(query, routeTmpId).Scan(&routeTmp.Id, &routeTmp.UserAuthorVkId,
		&routeTmp.LocDep, &routeTmp.LocArr, &routeTmp.MinPrice, &routeTmp.DateTimeDep,
		&routeTmp.DateTimeArr); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return routeTmp, nil
}

func (userRepository *UserRepository) InsertRoutePerm(routePerm *models.RoutePerm) (*models.RoutePerm, error) {
	const query = `
INSERT INTO view_route_perm (user_author_vk_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week, time_dep,
                             time_arr)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, user_author_vk_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week, time_dep, time_arr`

	if err := userRepository.db.QueryRow(query, routePerm.UserAuthorVkId, routePerm.LocDep, routePerm.LocArr,
		routePerm.MinPrice, routePerm.EvenWeek, routePerm.OddWeek, routePerm.DayOfWeek, time.Time(routePerm.TimeDep),
		time.Time(routePerm.TimeArr)).Scan(&routePerm.Id, &routePerm.UserAuthorVkId, &routePerm.LocDep,
		&routePerm.LocArr, &routePerm.MinPrice, &routePerm.EvenWeek, &routePerm.OddWeek, &routePerm.DayOfWeek,
		&routePerm.TimeDep, &routePerm.TimeArr); err != nil {
		return nil, err
	}

	return routePerm, nil
}
