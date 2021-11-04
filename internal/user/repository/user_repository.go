package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/user"
	"strconv"
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
	if err := userRepository.db.QueryRow(query, id).Scan(&user_.Id, &user_.VkId, &user_.Name, &user_.Avatar); err != nil {
		return nil, err
	}

	return user_, nil
}

func (userRepository *UserRepository) SelectByVkId(vkId uint32) (*models.User, error) {
	const query = "SELECT id, vk_id, name, avatar FROM user_ WHERE vk_id = $1"

	user_ := new(models.User)
	var avatar sql.NullString
	if err := userRepository.db.QueryRow(query, vkId).Scan(&user_.Id, &user_.VkId, &user_.Name, &avatar); err != nil {
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
	queryArgs := make([]interface{}, 1)
	queryArgs[0] = user_.Id

	if user_.Name != "" {
		query += queryName + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, user_.Name)
	}

	if user_.Avatar != "" {
		query += queryAvatar + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, user_.Avatar)
	}

	if len(queryArgs) == 1 {
		return userRepository.Select(user_.Id)
	}

	query = query[:len(query)-2] + queryEnd

	updatedUser := new(models.User)
	if err := userRepository.db.QueryRow(query, queryArgs...).Scan(&updatedUser.Id, &updatedUser.VkId,
		&updatedUser.Name, &updatedUser.Avatar); err != nil {
		return nil, err
	}

	return updatedUser, nil
}
