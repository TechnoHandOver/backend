package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/user/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_Insert(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	user := &models.User{
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	expectedUser := &models.User{
		Id:     1,
		VkId:   user.VkId,
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO user_").
		WithArgs(user.VkId, user.Name, user.Avatar).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "vk_id", "name", "avatar"}).
				AddRow(expectedUser.Id, user.VkId, user.Name, user.Avatar))

	resultUser, resultErr := userRepository.Insert(user)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedUser, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_Select(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	expectedUser := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}

	sqlmock_.
		ExpectQuery("SELECT id, vk_id, name, avatar FROM user_").
		WithArgs(expectedUser.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "vk_id", "name", "avatar"}).
				AddRow(expectedUser.Id, expectedUser.VkId, expectedUser.Name, expectedUser.Avatar))

	resultUser, resultErr := userRepository.Select(expectedUser.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedUser, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_Select_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const id uint32 = 1

	sqlmock_.
		ExpectQuery("SELECT id, vk_id, name, avatar FROM user_").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	resultUser, resultErr := userRepository.Select(id)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectByVkId(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	expectedUser := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}

	sqlmock_.
		ExpectQuery("SELECT id, vk_id, name, avatar FROM user_").
		WithArgs(expectedUser.VkId).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "vk_id", "name", "avatar"}).
				AddRow(expectedUser.Id, expectedUser.VkId, expectedUser.Name, expectedUser.Avatar))

	resultUser, resultErr := userRepository.SelectByVkId(expectedUser.VkId)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedUser, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectByVkId_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const vkId uint32 = 2

	sqlmock_.
		ExpectQuery("SELECT id, vk_id, name, avatar FROM user_").
		WithArgs(vkId).
		WillReturnError(sql.ErrNoRows)

	resultUser, resultErr := userRepository.SelectByVkId(vkId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	expectedUser := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}

	sqlmock_.
		ExpectQuery("UPDATE user_").
		WithArgs(expectedUser.Id, expectedUser.Name, expectedUser.Avatar).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "vk_id", "name", "avatar"}).
				AddRow(expectedUser.Id, expectedUser.VkId, expectedUser.Name, expectedUser.Avatar))

	resultUser, resultErr := userRepository.Update(expectedUser)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedUser, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_Update_select(t *testing.T) { //TODO: бизнес-логика там плохая, нужно переделать...
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	user := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "",
		Avatar: "",
	}
	expectedUser := &models.User{
		Id:     user.Id,
		VkId:   user.VkId,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}

	sqlmock_.
		ExpectQuery("SELECT id, vk_id, name, avatar FROM user_").
		WithArgs(user.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "vk_id", "name", "avatar"}).
				AddRow(expectedUser.Id, expectedUser.VkId, expectedUser.Name, expectedUser.Avatar))

	resultUser, resultErr := userRepository.Update(user)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedUser, resultUser)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_InsertRouteTmp(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep, err := timestamps.NewDateTime("10.11.2021 18:10")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("10.11.2021 18:15")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: routeTmp.UserAuthorId,
		LocDep:       routeTmp.LocDep,
		LocArr:       routeTmp.LocArr,
		MinPrice:     routeTmp.MinPrice,
		DateTimeDep:  routeTmp.DateTimeDep,
		DateTimeArr:  routeTmp.DateTimeArr,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO view_route_tmp").
		WithArgs(routeTmp.UserAuthorId, routeTmp.LocDep, routeTmp.LocArr, routeTmp.MinPrice,
			time.Time(routeTmp.DateTimeDep), time.Time(routeTmp.DateTimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, routeTmp.UserAuthorId, routeTmp.LocDep, routeTmp.LocArr,
					routeTmp.MinPrice, time.Time(routeTmp.DateTimeDep), time.Time(routeTmp.DateTimeArr)))

	resultRouteTmp, resultErr := userRepository.InsertRouteTmp(routeTmp)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRouteTmp, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRouteTmp(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 11:50")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 11:55")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp").
		WithArgs(expectedRouteTmp.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorId, expectedRouteTmp.LocDep,
					expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice, time.Time(expectedRouteTmp.DateTimeDep),
					time.Time(expectedRouteTmp.DateTimeArr)))

	resultRouteTmp, resultErr := userRepository.SelectRouteTmp(expectedRouteTmp.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRouteTmp, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRouteTmp_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const routeTmpId uint32 = 1

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp").
		WithArgs(routeTmpId).
		WillReturnError(sql.ErrNoRows)

	resultRouteTmp, resultErr := userRepository.SelectRouteTmp(routeTmpId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_UpdateRouteTmp(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 14:00")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 14:05")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("UPDATE view_route_tmp").
		WithArgs(expectedRouteTmp.Id, expectedRouteTmp.LocDep, expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice,
			time.Time(expectedRouteTmp.DateTimeDep), time.Time(expectedRouteTmp.DateTimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorId, expectedRouteTmp.LocDep,
					expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice, time.Time(expectedRouteTmp.DateTimeDep),
					time.Time(expectedRouteTmp.DateTimeArr)))

	resultRouteTmp, resultErr := userRepository.UpdateRouteTmp(expectedRouteTmp)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRouteTmp, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_UpdateRouteTmp_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 14:00")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 14:05")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("UPDATE view_route_tmp").
		WithArgs(routeTmp.Id, routeTmp.LocDep, routeTmp.LocArr, routeTmp.MinPrice, time.Time(routeTmp.DateTimeDep),
			time.Time(routeTmp.DateTimeArr)).
		WillReturnError(sql.ErrNoRows)

	resultRouteTmp, resultErr := userRepository.UpdateRouteTmp(routeTmp)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_DeleteRouteTmp(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 11:50")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 11:55")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("DELETE FROM view_route_tmp").
		WithArgs(expectedRouteTmp.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorId, expectedRouteTmp.LocDep,
					expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice, time.Time(expectedRouteTmp.DateTimeDep),
					time.Time(expectedRouteTmp.DateTimeArr)))

	resultRouteTmp, resultErr := userRepository.DeleteRouteTmp(expectedRouteTmp.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRouteTmp, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_DeleteRouteTmp_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const routeTmpId uint32 = 1

	sqlmock_.
		ExpectQuery("DELETE FROM view_route_tmp").
		WithArgs(routeTmpId).
		WillReturnError(sql.ErrNoRows)

	resultRouteTmp, resultErr := userRepository.DeleteRouteTmp(routeTmpId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRouteTmpArray(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	dateTimeDep1, err := timestamps.NewDateTime("17.11.2021 10:25")
	assert.Nil(t, err)
	dateTimeArr1, err := timestamps.NewDateTime("17.11.2021 10:30")
	assert.Nil(t, err)
	dateTimeDep2, err := timestamps.NewDateTime("17.11.2021 10:35")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("17.11.2021 10:40")
	assert.Nil(t, err)
	expectedRoutesTmp := &models.RoutesTmp{
		&models.RouteTmp{
			Id:           1,
			UserAuthorId: 101,
			LocDep:       "Общежитие №10",
			LocArr:       "УЛК",
			MinPrice:     500,
			DateTimeDep:  *dateTimeDep1,
			DateTimeArr:  *dateTimeArr1,
		},
		&models.RouteTmp{
			Id:           2,
			UserAuthorId: 102,
			LocDep:       "Общежитие №9",
			LocArr:       "СК",
			MinPrice:     600,
			DateTimeDep:  *dateTimeDep2,
			DateTimeArr:  *dateTimeArr2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
		"date_time_arr"})
	for _, expectedRouteTmp := range *expectedRoutesTmp {
		rows.AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorId, expectedRouteTmp.LocDep,
			expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice, time.Time(expectedRouteTmp.DateTimeDep),
			time.Time(expectedRouteTmp.DateTimeArr))
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp").
		WillReturnRows(rows)

	resultRoutesTmp, resultErr := userRepository.SelectRouteTmpArray()
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutesTmp, resultRoutesTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_InsertRoutePerm(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep, err := timestamps.NewTime("12:30")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("12:35")
	assert.Nil(t, err)
	routePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}
	expectedRoutePerm := &models.RoutePerm{
		Id:           1,
		UserAuthorId: routePerm.UserAuthorId,
		LocDep:       routePerm.LocDep,
		LocArr:       routePerm.LocArr,
		MinPrice:     routePerm.MinPrice,
		EvenWeek:     routePerm.EvenWeek,
		OddWeek:      routePerm.OddWeek,
		DayOfWeek:    routePerm.DayOfWeek,
		TimeDep:      routePerm.TimeDep,
		TimeArr:      routePerm.TimeArr,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO view_route_perm").
		WithArgs(routePerm.UserAuthorId, routePerm.LocDep, routePerm.LocArr, routePerm.MinPrice, routePerm.EvenWeek,
			routePerm.OddWeek, routePerm.DayOfWeek, time.Time(routePerm.TimeDep), time.Time(routePerm.TimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "even_week",
				"odd_week", "day_of_week", "time_dep", "time_arr"}).
				AddRow(expectedRoutePerm.Id, routePerm.UserAuthorId, routePerm.LocDep, routePerm.LocArr,
					routePerm.MinPrice, routePerm.EvenWeek, routePerm.OddWeek, routePerm.DayOfWeek,
					time.Time(routePerm.TimeDep), time.Time(routePerm.TimeArr)))

	resultRoutePerm, resultErr := userRepository.InsertRoutePerm(routePerm)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutePerm, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRoutePerm(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week, time_dep, time_arr FROM view_route_perm").
		WithArgs(expectedRoutePerm.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "even_week",
				"odd_week", "day_of_week", "time_dep", "time_arr"}).
				AddRow(expectedRoutePerm.Id, expectedRoutePerm.UserAuthorId, expectedRoutePerm.LocDep,
					expectedRoutePerm.LocArr, expectedRoutePerm.MinPrice, expectedRoutePerm.EvenWeek,
					expectedRoutePerm.OddWeek, expectedRoutePerm.DayOfWeek, time.Time(expectedRoutePerm.TimeDep),
					time.Time(expectedRoutePerm.TimeArr)))

	resultRoutePerm, resultErr := userRepository.SelectRoutePerm(expectedRoutePerm.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutePerm, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRoutePerm_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const routePermId uint32 = 1

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week, time_dep, time_arr FROM view_route_perm").
		WithArgs(routePermId).
		WillReturnError(sql.ErrNoRows)

	resultRoutePerm, resultErr := userRepository.SelectRoutePerm(routePermId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_UpdateRoutePerm(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	sqlmock_.
		ExpectQuery("UPDATE view_route_perm").
		WithArgs(expectedRoutePerm.Id, expectedRoutePerm.LocDep, expectedRoutePerm.LocArr, expectedRoutePerm.MinPrice,
			expectedRoutePerm.EvenWeek, expectedRoutePerm.OddWeek, expectedRoutePerm.DayOfWeek,
			time.Time(expectedRoutePerm.TimeDep), time.Time(expectedRoutePerm.TimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "even_week",
				"odd_week", "day_of_week", "time_dep", "time_arr"}).
				AddRow(expectedRoutePerm.Id, expectedRoutePerm.UserAuthorId, expectedRoutePerm.LocDep,
					expectedRoutePerm.LocArr, expectedRoutePerm.MinPrice, expectedRoutePerm.EvenWeek,
					expectedRoutePerm.OddWeek, expectedRoutePerm.DayOfWeek, time.Time(expectedRoutePerm.TimeDep),
					time.Time(expectedRoutePerm.TimeArr)))

	resultRoutePerm, resultErr := userRepository.UpdateRoutePerm(expectedRoutePerm)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutePerm, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_UpdateRoutePerm_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	routePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	sqlmock_.
		ExpectQuery("UPDATE view_route_perm").
		WithArgs(routePerm.Id, routePerm.LocDep, routePerm.LocArr, routePerm.MinPrice, routePerm.EvenWeek,
			routePerm.OddWeek, routePerm.DayOfWeek, time.Time(routePerm.TimeDep), time.Time(routePerm.TimeArr)).
		WillReturnError(sql.ErrNoRows)

	resultRouteTmp, resultErr := userRepository.UpdateRoutePerm(routePerm)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRouteTmp)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_DeleteRoutePerm(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	sqlmock_.
		ExpectQuery("DELETE FROM view_route_perm").
		WithArgs(expectedRoutePerm.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "even_week",
				"odd_week", "day_of_week", "time_dep", "time_arr"}).
				AddRow(expectedRoutePerm.Id, expectedRoutePerm.UserAuthorId, expectedRoutePerm.LocDep,
					expectedRoutePerm.LocArr, expectedRoutePerm.MinPrice, expectedRoutePerm.EvenWeek,
					expectedRoutePerm.OddWeek, expectedRoutePerm.DayOfWeek, time.Time(expectedRoutePerm.TimeDep),
					time.Time(expectedRoutePerm.TimeArr)))

	resultRoutePerm, resultErr := userRepository.DeleteRoutePerm(expectedRoutePerm.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutePerm, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_DeleteRoutePerm_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	const routePermId uint32 = 1

	sqlmock_.
		ExpectQuery("DELETE FROM view_route_perm").
		WithArgs(routePermId).
		WillReturnError(sql.ErrNoRows)

	resultRoutePerm, resultErr := userRepository.DeleteRoutePerm(routePermId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultRoutePerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestUserRepository_SelectRoutePermArray(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	timeDep1, err := timestamps.NewTime("17:30")
	assert.Nil(t, err)
	timeArr1, err := timestamps.NewTime("17:35")
	assert.Nil(t, err)
	timeDep2, err := timestamps.NewTime("17:40")
	assert.Nil(t, err)
	timeArr2, err := timestamps.NewTime("17:45")
	assert.Nil(t, err)
	expectedRoutesPerm := &models.RoutesPerm{
		&models.RoutePerm{
			Id:           1,
			UserAuthorId: 101,
			LocDep:       "Общежитие №10",
			LocArr:       "УЛК",
			MinPrice:     500,
			EvenWeek:     true,
			OddWeek:      false,
			DayOfWeek:    timestamps.DayOfWeekWednesday,
			TimeDep:      *timeDep1,
			TimeArr:      *timeArr1,
		},
		&models.RoutePerm{
			Id:           1,
			UserAuthorId: 102,
			LocDep:       "Общежитие №9",
			LocArr:       "СК",
			MinPrice:     600,
			EvenWeek:     false,
			OddWeek:      true,
			DayOfWeek:    timestamps.DayOfWeekSaturday,
			TimeDep:      *timeDep2,
			TimeArr:      *timeArr2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_id", "loc_dep", "loc_arr", "min_price", "even_week",
		"odd_week", "day_of_week", "time_dep", "time_arr"})
	for _, expectedRoutePerm := range *expectedRoutesPerm {
		rows.AddRow(expectedRoutePerm.Id, expectedRoutePerm.UserAuthorId, expectedRoutePerm.LocDep,
			expectedRoutePerm.LocArr, expectedRoutePerm.MinPrice, expectedRoutePerm.EvenWeek,
			expectedRoutePerm.OddWeek, expectedRoutePerm.DayOfWeek, time.Time(expectedRoutePerm.TimeDep),
			time.Time(expectedRoutePerm.TimeArr))
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week, time_dep, time_arr FROM view_route_perm").
		WillReturnRows(rows)

	resultRoutesPerm, resultErr := userRepository.SelectRoutePermArray()
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedRoutesPerm, resultRoutesPerm)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}
