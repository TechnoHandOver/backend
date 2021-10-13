package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"strconv"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) Insert(user *models.User) (*models.User, error) {
	const query = "INSERT INTO user_ (vk_id, name, avatar) VALUES ($1, $2, $3) RETURNING id, vk_id, name, avatar"

	if err := userRepository.db.QueryRow(query, user.VkId, user.Name, user.Avatar).Scan(&user.Id, &user.VkId,
		&user.Name, &user.Avatar); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) Select(id uint32) (*models.User, error) {
	const query = "SELECT id, vk_id, name, avatar FROM user_ WHERE id = $1"

	user := new(models.User)
	if err := userRepository.db.QueryRow(query, id).Scan(&user.Id, &user.VkId, &user.Name, &user.Avatar); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) SelectByVkId(vkId uint32) (*models.User, error) {
	const query = "SELECT id, vk_id, name, avatar FROM user_ WHERE vk_id = $1"

	user := new(models.User)
	if err := userRepository.db.QueryRow(query, vkId).Scan(&user.Id, &user.VkId, &user.Name, &user.Avatar); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) Update(id uint32, userUpdate *models.UserUpdate) (*models.User, error) {
	const queryStart = "UPDATE user_ SET "
	const queryName = "name"
	const queryAvatar = "avatar"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = "WHERE id = $1 RETURNING id, vk_id, name, avatar"

	query := queryStart
	queryArgs := make([]interface{}, 1)
	queryArgs[0] = strconv.FormatUint(uint64(id), 10)

	if userUpdate.Name != "" {
		query += queryName + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, userUpdate.Name)
	}

	if userUpdate.Avatar != "" {
		query += queryAvatar + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, userUpdate.Avatar)
	}

	if len(queryArgs) == 1 {
		return userRepository.Select(id)
	}

	query = query[:len(query)-2] + queryEnd

	updatedUser := new(models.User)
	if err := userRepository.db.QueryRow(query, queryArgs...).Scan(&updatedUser.Id, &updatedUser.VkId,
		&updatedUser.Name, &updatedUser.Avatar); err != nil {
		return nil, err
	}

	return updatedUser, nil
}
