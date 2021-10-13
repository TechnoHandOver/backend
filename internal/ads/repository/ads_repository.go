package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
)

type AdsRepository struct {
	db *sql.DB
}

func NewAdsRepository(db *sql.DB) *AdsRepository {
	return &AdsRepository{
		db: db,
	}
}

func (adsRepository *AdsRepository) Insert(ads *models.Ads) (*models.Ads, error) {
	const query = "INSERT INTO ads (user_author_id, location_from, location_to, time_from, time_to, min_price, comment) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING ads.id, ads.user_author_id, ads.location_from, ads.location_to, ads.time_from, ads.time_to, ads.min_price, ads.comment"

	if err := adsRepository.db.QueryRow(query, ads.UserAuthorId, ads.LocationFrom, ads.LocationTo, ads.TimeFrom,
		ads.TimeTo, ads.MinPrice, ads.Comment).Scan(&ads.Id, &ads.UserAuthorId, &ads.LocationFrom, &ads.LocationTo,
			&ads.TimeFrom, &ads.TimeTo, &ads.MinPrice, &ads.Comment); err != nil {
		return nil, err
	}

	return ads, nil
}

func (adsRepository *AdsRepository) Select(id uint32) (*models.Ads, error) {
	const query = "SELECT id, user_author_id, location_from, location_to, time_from, time_to, min_price, comment FROM ads WHERE id = $1"

	ads := new(models.Ads)
	if err := adsRepository.db.QueryRow(query, id).Scan(&ads.Id, &ads.UserAuthorId, &ads.LocationFrom,
		&ads.LocationTo, &ads.TimeFrom, &ads.TimeTo, &ads.MinPrice, &ads.Comment); err != nil {
		return nil, err
	}

	return ads, nil
}

func (adsRepository *AdsRepository) SelectArray() (*models.Adses, error) {
	const query = "SELECT id, user_author_id, location_from, location_to, time_from, time_to, min_price, comment FROM ads ORDER BY time_to"

	rows, err := adsRepository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	adses := make(models.Adses, 0)
	for rows.Next() {
		ads := new(models.Ads)
		if err := rows.Scan(&ads.Id, &ads.UserAuthorId, &ads.LocationFrom, &ads.LocationTo, &ads.TimeFrom, &ads.TimeTo,
			&ads.MinPrice, &ads.Comment); err != nil {
			return nil, err
		}

		adses = append(adses, ads)
	}

	return &adses, nil
}
