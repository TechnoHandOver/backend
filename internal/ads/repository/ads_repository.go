package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"strconv"
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

func (adsRepository *AdsRepository) Update(id uint32, adsUpdate *models.AdsUpdate) (*models.Ads, error) {
	const queryStart = "UPDATE ads SET "
	const queryLocationFrom = "location_from"
	const queryLocationTo = "location_to"
	const queryTimeFrom = "time_from"
	const queryTimeTo = "time_to"
	const queryMinPrice = "min_price"
	const queryComment = "comment"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = "WHERE id = $1 RETURNING ads.id, ads.user_author_id, ads.location_from, ads.location_to, ads.time_from, ads.time_to, ads.min_price, ads.comment"

	query := queryStart
	queryArgs := make([]interface{}, 1)
	queryArgs[0] = strconv.FormatUint(uint64(id), 10)

	if adsUpdate.LocationFrom != "" {
		query += queryLocationFrom + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.LocationFrom)
	}

	if adsUpdate.LocationTo != "" {
		query += queryLocationTo + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.LocationTo)
	}

	if !adsUpdate.TimeFrom.IsZero() {
		query += queryTimeFrom + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.TimeFrom)
	}

	if !adsUpdate.TimeTo.IsZero() {
		query += queryTimeTo + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.TimeTo)
	}

	if adsUpdate.MinPrice != 0 {
		query += queryMinPrice + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.MinPrice)
	}

	if adsUpdate.Comment != "" {
		query += queryComment + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.Comment)
	}

	if len(queryArgs) == 1 {
		return adsRepository.Select(id)
	}

	query = query[:len(query)-2] + queryEnd

	updatedAds := new(models.Ads)
	if err := adsRepository.db.QueryRow(query, queryArgs...).Scan(&updatedAds.Id, &updatedAds.UserAuthorId,
		&updatedAds.LocationFrom, &updatedAds.LocationTo, &updatedAds.TimeFrom, &updatedAds.TimeTo,
		&updatedAds.MinPrice, &updatedAds.Comment); err != nil {
		return nil, err
	}

	return updatedAds, nil
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
