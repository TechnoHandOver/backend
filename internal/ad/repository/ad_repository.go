package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/models"
	"strconv"
	"time"
)

type AdRepository struct {
	db *sql.DB
}

func NewAdRepositoryImpl(db *sql.DB) ad.Repository {
	return &AdRepository{
		db: db,
	}
}

func (adsRepository *AdRepository) Insert(ad_ *models.Ad) (*models.Ad, error) {
	const query = `
INSERT INTO ad (user_author_id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment)
SELECT user_.id, user_.vk_id, $2, $3, $4, $5, $6, $7 FROM user_ WHERE user_.vk_id = $1
RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment`

	var dateTimeArr = time.Time(ad_.DateTimeArr)
	if err := adsRepository.db.QueryRow(query, ad_.UserAuthorVkId, ad_.LocDep, ad_.LocArr, dateTimeArr, ad_.Item,
		ad_.MinPrice, ad_.Comment).Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr, &ad_.DateTimeArr,
		&ad_.Item, &ad_.MinPrice, &ad_.Comment); err != nil {
		return nil, err
	}

	return ad_, nil
}

func (adsRepository *AdRepository) Select(id uint32) (*models.Ad, error) {
	const query = `
SELECT id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment
FROM ad
WHERE id = $1`

	ad_ := new(models.Ad)
	if err := adsRepository.db.QueryRow(query, id).Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr,
		&ad_.DateTimeArr, &ad_.Item, &ad_.MinPrice, &ad_.Comment); err != nil {
		return nil, err
	}

	return ad_, nil
}

func (adsRepository *AdRepository) Update(ad_ *models.Ad) (*models.Ad, error) {
	const queryStart = "UPDATE ad SET "
	const queryLocDep = "loc_dep"
	const queryLocArr = "loc_arr"
	const queryDateArr = "date_time_arr"
	const queryItem = "item"
	const queryMinPrice = "min_price"
	const queryComment = "comment"
	const queryEquals = " = $"
	const queryComma = ", "
	const queryEnd = "WHERE id = $1 RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment"

	query := queryStart
	queryArgs := make([]interface{}, 1)
	queryArgs[0] = ad_.Id

	if ad_.LocDep != "" {
		query += queryLocDep + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, ad_.LocDep)
	}

	if ad_.LocArr != "" {
		query += queryLocArr + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, ad_.LocArr)
	}

	if dateArrTime := time.Time(ad_.DateTimeArr); !dateArrTime.IsZero() {
		query += queryDateArr + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, dateArrTime)
	}

	if ad_.Item != "" {
		query += queryItem + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, ad_.Item)
	}

	if ad_.MinPrice != 0 {
		query += queryMinPrice + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, ad_.MinPrice)
	}

	if ad_.Comment != "" {
		query += queryComment + queryEquals + strconv.Itoa(len(queryArgs)+1) + queryComma
		queryArgs = append(queryArgs, ad_.Comment)
	}

	if len(queryArgs) == 1 {
		return adsRepository.Select(ad_.Id) //TODO: возможно, нарушение логики и зон ответственности...; не только здесь так
	}

	query = query[:len(query)-2] + queryEnd

	updatedAd := new(models.Ad)
	if err := adsRepository.db.QueryRow(query, queryArgs...).Scan(&updatedAd.Id, &updatedAd.UserAuthorVkId,
		&updatedAd.LocDep, &updatedAd.LocArr, &updatedAd.DateTimeArr, &updatedAd.Item, &updatedAd.MinPrice,
		&updatedAd.Comment); err != nil {
		return nil, err
	}

	return updatedAd, nil
}

func (adsRepository *AdRepository) SelectArray(adsSearch *models.AdsSearch) (*models.Ads, error) {
	const queryStart = "SELECT id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad"
	const queryWhere = " WHERE "
	const queryLocDep1 = "to_tsvector('russian', loc_dep) @@ plainto_tsquery('russian', $"
	const queryLocDep2 = ")"
	const queryLocArr1 = "to_tsvector('russian', loc_arr) @@ plainto_tsquery('russian', $"
	const queryLocArr2 = ")"
	const queryDateArr = "date_time_arr = $"
	const queryMinPrice = "min_price <= $"
	const queryAnd = " AND "
	const queryEnd = " ORDER BY min_price"

	query := queryStart + queryWhere
	queryArgs := make([]interface{}, 0)

	if adsSearch.LocDep != "" {
		query += queryLocDep1 + strconv.Itoa(len(queryArgs)+1) + queryLocDep2 + queryAnd
		queryArgs = append(queryArgs, adsSearch.LocDep)
	}

	if adsSearch.LocArr != "" {
		query += queryLocArr1 + strconv.Itoa(len(queryArgs)+1) + queryLocArr2 + queryAnd
		queryArgs = append(queryArgs, adsSearch.LocArr)
	}

	if dateTimeArrTime := time.Time(adsSearch.DateTimeArr); !dateTimeArrTime.IsZero() {
		query += queryDateArr + strconv.Itoa(len(queryArgs)+1) + queryAnd
		queryArgs = append(queryArgs, dateTimeArrTime)
	}

	if adsSearch.MaxPrice != 0 {
		query += queryMinPrice + strconv.Itoa(len(queryArgs)+1) + queryAnd
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

	ads := make(models.Ads, 0)
	for rows.Next() {
		ad_ := new(models.Ad)
		if err := rows.Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr, &ad_.DateTimeArr, &ad_.Item,
			&ad_.MinPrice, &ad_.Comment); err != nil {
			return nil, err
		}

		ads = append(ads, ad_)
	}

	return &ads, nil
}
