package repository

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/consts"
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

	if err := adsRepository.db.QueryRow(query, ad_.UserAuthorVkId, ad_.LocDep, ad_.LocArr, time.Time(ad_.DateTimeArr),
		ad_.Item, ad_.MinPrice, ad_.Comment).Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr,
		&ad_.DateTimeArr, &ad_.Item, &ad_.MinPrice, &ad_.Comment); err != nil {
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
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return ad_, nil
}

func (adsRepository *AdRepository) Update(ad_ *models.Ad) (*models.Ad, error) {
	const query = `
UPDATE ad SET loc_dep = $2, loc_arr = $3, date_time_arr = $4, item = $5, min_price = $6, comment = $7
WHERE id = $1
RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment`

	if err := adsRepository.db.QueryRow(query, ad_.Id, ad_.LocDep, ad_.LocArr, time.Time(ad_.DateTimeArr), ad_.Item,
		ad_.MinPrice, ad_.Comment).Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr, &ad_.DateTimeArr,
		&ad_.Item, &ad_.MinPrice, &ad_.Comment); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return ad_, nil
}

func (adsRepository *AdRepository) Delete(id uint32) (*models.Ad, error) {
	const query = `
DELETE FROM ad
WHERE id = $1
RETURNING id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment`

	ad_ := new(models.Ad)
	if err := adsRepository.db.QueryRow(query, id).Scan(&ad_.Id, &ad_.UserAuthorVkId, &ad_.LocDep, &ad_.LocArr,
		&ad_.DateTimeArr, &ad_.Item, &ad_.MinPrice, &ad_.Comment); err != nil {
		if err == sql.ErrNoRows {
			return nil, consts.RepErrNotFound
		}

		return nil, err
	}

	return ad_, nil
}

func (adsRepository *AdRepository) SelectArray(adsSearch *models.AdsSearch) (*models.Ads, error) {
	const queryStart = "SELECT id, user_author_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad"
	const queryWhere = " WHERE "
	const queryLocDep1 = "to_tsvector('russian', loc_dep) @@ plainto_tsquery('russian', $"
	const queryLocDep2 = ")"
	const queryLocArr1 = "to_tsvector('russian', loc_arr) @@ plainto_tsquery('russian', $"
	const queryLocArr2 = ")"
	const queryDateTimeArr = "date_time_arr = $"
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

	if dateTimeArr := time.Time(adsSearch.DateTimeArr); !dateTimeArr.IsZero() {
		query += queryDateTimeArr + strconv.Itoa(len(queryArgs)+1) + queryAnd
		queryArgs = append(queryArgs, dateTimeArr)
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
