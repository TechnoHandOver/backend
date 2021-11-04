package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/user/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Insert(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	user := &models.User{
		VkId: 2,
		Name: "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	expectedUser := &models.User {
		Id: 1,
		VkId: user.VkId,
		Name: user.Name,
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
		Id: 1,
		VkId: 2,
		Name: "Василий Петров",
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

func TestUserRepository_SelectByVkId(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	expectedUser := &models.User{
		Id: 1,
		VkId: 2,
		Name: "Василий Петров",
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

func TestUserRepository_Update(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	expectedUser := &models.User{
		Id: 1,
		VkId: 2,
		Name: "Василий Петров",
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

func TestUserRepository_Update_select(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	userRepository := repository.NewUserRepositoryImpl(db)

	user := &models.User{
		Id: 1,
		VkId: 2,
		Name: "",
		Avatar: "",
	}
	expectedUser := &models.User{
		Id: user.Id,
		VkId: user.VkId,
		Name: "Василий Петров",
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
