package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TechnoHandOver/backend/internal/ad/repository"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdRepository_Insert(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
	assert.Nil(t, err)
	ad := &models.Ad{
		UserAuthorVkId: 2,
		LocDep: "Общежитие №10",
		LocArr: "УЛК",
		DateTimeArr: *dateTimeArr,
		Item: "Зачётная книжка",
		MinPrice: 500,
		Comment: "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id: 1,
		UserAuthorVkId: ad.UserAuthorVkId,
		LocDep: ad.LocDep,
		LocArr: ad.LocArr,
		DateTimeArr: ad.DateTimeArr,
		Item: ad.Item,
		MinPrice: ad.MinPrice,
		Comment: ad.Comment,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO ad").
		WithArgs(ad.UserAuthorVkId, ad.LocDep, ad.LocArr, time.Time(ad.DateTimeArr), ad.Item, ad.MinPrice, ad.Comment).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price", "comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Insert(ad)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Select(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id: 1,
		UserAuthorVkId: 2,
		LocDep: "Общежитие №10",
		LocArr: "УЛК",
		DateTimeArr: *dateTimeArr,
		Item: "Зачётная книжка",
		MinPrice: 500,
		Comment: "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, item, min_price, comment FROM ad").
		WithArgs(expectedAd.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price", "comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Select(expectedAd.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Update(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id: 1,
		UserAuthorVkId: 2,
		LocDep: "Общежитие №10",
		LocArr: "УЛК",
		DateTimeArr: *dateTimeArr,
		Item: "Зачётная книжка",
		MinPrice: 500,
		Comment: "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("UPDATE ad").
		WithArgs(expectedAd.Id, expectedAd.LocDep, expectedAd.LocArr, time.Time(expectedAd.DateTimeArr),
			expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price", "comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Update(expectedAd)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Update_select(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	ad := &models.Ad{
		Id: 1,
		UserAuthorVkId: 2,
	}
	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id: ad.Id,
		UserAuthorVkId: ad.UserAuthorVkId,
		LocDep: "Общежитие №10",
		LocArr: "УЛК",
		DateTimeArr: *dateTimeArr,
		Item: "Зачётная книжка",
		MinPrice: 500,
		Comment: "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, item, min_price, comment FROM ad").
		WithArgs(expectedAd.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price", "comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Update(ad)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_SelectArray(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := &models.AdsSearch{
		LocDep: "Общежитие",
		LocArr: "СК",
		DateTimeArr: *dateTimeArr1,
		MaxPrice: 1000,
	}
	expectedAds := &models.Ads{
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №10",
			LocArr: "УЛК",
			DateTimeArr: *dateTimeArr1,
			Item: "Тубус",
			MinPrice: 500,
			Comment: "Поеду на коньках",
		},
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №9",
			LocArr: "СК",
			DateTimeArr: *dateTimeArr2,
			Item: "Спортивная форма",
			MinPrice: 600,
			Comment: "Поеду на роликах :)",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price", "comment"})
	for _, expectedAd := range *expectedAds {
		rows.AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
			time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment)
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, item, min_price, comment FROM ad").
		WithArgs(adsSearch.LocDep, adsSearch.LocArr, time.Time(adsSearch.DateTimeArr), adsSearch.MaxPrice).
		WillReturnRows(rows)

	resultAds, resultErr := adRepository.SelectArray(adsSearch)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAds, resultAds)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_SelectArray_emptySearchQuery(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := new(models.AdsSearch)
	expectedAds := &models.Ads{
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №10",
			LocArr: "СК",
			DateTimeArr: *dateTimeArr1,
			MinPrice: 500,
			Comment: "Поеду на коньках",
		},
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №9",
			LocArr: "СК",
			DateTimeArr: *dateTimeArr2,
			MinPrice: 600,
			Comment: "Поеду на роликах :)",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_vk_id", "loc_dep", "loc_dep", "date_arr", "item", "min_price",
		"comment"})
	for _, expectedAd := range *expectedAds {
		rows.AddRow(expectedAd.Id, expectedAd.UserAuthorVkId, expectedAd.LocDep, expectedAd.LocArr,
			time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment)
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_vk_id, loc_dep, loc_arr, date_arr, item, min_price, comment FROM ad").
		WillReturnRows(rows)

	resultAds, resultErr := adRepository.SelectArray(adsSearch)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAds, resultAds)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}
