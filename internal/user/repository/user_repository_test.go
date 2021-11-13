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
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}
	expectedRouteTmp := &models.RouteTmp{
		Id:             1,
		UserAuthorVkId: routeTmp.UserAuthorVkId,
		LocDep:         routeTmp.LocDep,
		LocArr:         routeTmp.LocArr,
		MinPrice:       routeTmp.MinPrice,
		DateTimeDep:    routeTmp.DateTimeDep,
		DateTimeArr:    routeTmp.DateTimeArr,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO view_route_tmp").
		WithArgs(routeTmp.UserAuthorVkId, routeTmp.LocDep, routeTmp.LocArr, routeTmp.MinPrice,
			time.Time(routeTmp.DateTimeDep), time.Time(routeTmp.DateTimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, routeTmp.UserAuthorVkId, routeTmp.LocDep, routeTmp.LocArr,
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
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp").
		WithArgs(expectedRouteTmp.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorVkId, expectedRouteTmp.LocDep,
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
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr FROM view_route_tmp").
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
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("UPDATE view_route_tmp").
		WithArgs(expectedRouteTmp.Id, expectedRouteTmp.LocDep, expectedRouteTmp.LocArr, expectedRouteTmp.MinPrice,
			time.Time(expectedRouteTmp.DateTimeDep), time.Time(expectedRouteTmp.DateTimeArr)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorVkId, expectedRouteTmp.LocDep,
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
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
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
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}

	sqlmock_.
		ExpectQuery("DELETE FROM view_route_tmp").
		WithArgs(expectedRouteTmp.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_arr", "min_price", "date_time_dep",
				"date_time_arr"}).
				AddRow(expectedRouteTmp.Id, expectedRouteTmp.UserAuthorVkId, expectedRouteTmp.LocDep,
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
