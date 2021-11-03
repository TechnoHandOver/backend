package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"strconv"
	"time"
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
	const query = `
INSERT INTO ads (user_author_id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment)
SELECT user_.id, user_.vk_id, $2, $3, $4, $5, $6 FROM user_ WHERE user_.vk_id = $1
RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment`

	var dateArrTime = time.Time(ads.DateTimeArr)
	if err := adsRepository.db.QueryRow(query, ads.UserAuthorVkId, ads.LocDep, ads.LocArr, dateArrTime, ads.MinPrice,
		ads.Comment).Scan(&ads.Id, &ads.UserAuthorVkId, &ads.LocDep, &ads.LocArr, &dateArrTime, &ads.MinPrice,
			&ads.Comment); err != nil {
		return nil, err
	}

	return ads, nil
}

func (adsRepository *AdsRepository) Select(id uint32) (*models.Ads, error) {
	const query = `
SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment
FROM ads
WHERE id = $1`

	ads := new(models.Ads)
	if err := adsRepository.db.QueryRow(query, id).Scan(&ads.Id, &ads.UserAuthorVkId, &ads.LocDep, &ads.LocArr,
		&ads.DateTimeArr, &ads.MinPrice, &ads.Comment); err != nil {
		return nil, err
	}

	return ads, nil
}

func (adsRepository *AdsRepository) Update(id uint32, adsUpdate *models.AdsUpdate) (*models.Ads, error) {
	const queryStart = "UPDATE ads SET "
	const queryLocDep = "loc_dep"
	const queryLocArr = "loc_arr"
	const queryDateArr = "date_arr"
	const queryMinPrice = "min_price"
	const queryComment = "comment"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = "WHERE id = $1 RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment"

	query := queryStart
	queryArgs := make([]interface{}, 1)
	queryArgs[0] = strconv.FormatUint(uint64(id), 10)

	if adsUpdate.LocDep != "" {
		query += queryLocDep + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.LocDep)
	}

	if adsUpdate.LocArr != "" {
		query += queryLocArr + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, adsUpdate.LocArr)
	}

	if dateArrTime := time.Time(adsUpdate.DateTimeArr); !dateArrTime.IsZero() {
		query += queryDateArr + queryEquals + strconv.Itoa(len(queryArgs) + 1) + queryComma
		queryArgs = append(queryArgs, dateArrTime)
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
	if err := adsRepository.db.QueryRow(query, queryArgs...).Scan(&updatedAds.Id, &updatedAds.UserAuthorVkId,
		&updatedAds.LocDep, &updatedAds.LocArr, &updatedAds.DateTimeArr, &updatedAds.MinPrice,
		&updatedAds.Comment); err != nil {
		return nil, err
	}

	return updatedAds, nil
}

func (adsRepository *AdsRepository) SelectArray() (*models.Adses, error) {
	const query = `
SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment
FROM ads
WHERE id = $1
ORDER BY date_arr`

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
		if err := rows.Scan(&ads.Id, &ads.UserAuthorVkId, &ads.LocDep, &ads.LocArr, &ads.DateTimeArr, &ads.MinPrice,
			&ads.Comment); err != nil {
			return nil, err
		}

		adses = append(adses, ads)
	}

	return &adses, nil
}

// SelectArray2 TODO: переименовать метод после удаления запроса на список объявлений
func (adsRepository *AdsRepository) SelectArray2(adsSearch *models.AdsSearch) (*models.Adses, error) {
	const queryStart = "SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, min_price, comment FROM ads"
	const queryWhere = " WHERE "
	const queryLocDep1 = "to_tsvector('russian', loc_dep) @@ plainto_tsquery('russian', $"
	const queryLocDep2 = ")"
	const queryLocArr1 = "to_tsvector('russian', loc_arr) @@ plainto_tsquery('russian', $"
	const queryLocArr2 = ")"
	const queryDateArr = "date_arr = $"
	const queryMinPrice = "min_price <= $"
	const queryAnd = " AND "
	const queryEnd = " ORDER BY min_price"

	query := queryStart + queryWhere
	queryArgs := make([]interface{}, 0)

	if adsSearch.LocDep != "" {
		query += queryLocDep1 + strconv.Itoa(len(queryArgs) + 1) + queryLocDep2 + queryAnd
		queryArgs = append(queryArgs, adsSearch.LocDep)
	}

	if adsSearch.LocArr != "" {
		query += queryLocArr1 + strconv.Itoa(len(queryArgs) + 1) + queryLocArr2 + queryAnd
		queryArgs = append(queryArgs, adsSearch.LocArr)
	}

	if dateTimeArrTime := time.Time(adsSearch.DateTimeArr); !dateTimeArrTime.IsZero() {
		query += queryDateArr + strconv.Itoa(len(queryArgs) + 1) + queryAnd
		queryArgs = append(queryArgs, dateTimeArrTime)
	}

	if adsSearch.MaxPrice != 0 {
		query += queryMinPrice + strconv.Itoa(len(queryArgs) + 1) + queryAnd
		queryArgs = append(queryArgs, adsSearch.MaxPrice)
	}

	if len(queryArgs) == 0 {
		query = query[:len(query)-len(queryWhere)]
	} else {
		query = query[:len(query)-len(queryAnd)]
	}
	query += queryEnd

	rows, err := adsRepository.db.Query(query, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	adses := make(models.Adses, 0)
	for rows.Next() {
		ads := new(models.Ads)
		if err := rows.Scan(&ads.Id, &ads.UserAuthorVkId, &ads.LocDep, &ads.LocArr, &ads.DateTimeArr, &ads.MinPrice,
			&ads.Comment); err != nil {
			return nil, err
		}

		adses = append(adses, ads)
	}

	return &adses, nil
}
