package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TechnoHandOver/backend/internal/ad/repository"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	HandoverTesting "github.com/TechnoHandOver/backend/internal/tools/testing"
	"github.com/openlyinc/pointy"
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
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:               1,
		UserAuthorId:     ad.UserAuthorId,
		UserAuthorVkId:   201,
		UserAuthorName:   "Vasiliy Pupkin",
		UserAuthorAvatar: "https://yandex.ru/logo.png",
		LocDep:           ad.LocDep,
		LocArr:           ad.LocArr,
		DateTimeArr:      ad.DateTimeArr,
		Item:             ad.Item,
		MinPrice:         ad.MinPrice,
		Comment:          ad.Comment,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO ad").
		WithArgs(ad.UserAuthorId, ad.LocDep, ad.LocArr, time.Time(ad.DateTimeArr), ad.Item, ad.MinPrice, ad.Comment).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
				"user_author_avatar", "loc_dep", "loc_dep", "date_time_arr", "item", "min_price", "comment"}).
				AddRow(expectedAd.Id, ad.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
					expectedAd.UserAuthorAvatar, ad.LocDep, ad.LocArr, time.Time(ad.DateTimeArr), ad.Item, ad.MinPrice,
					ad.Comment))

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
		Id:               1,
		UserAuthorId:     101,
		UserAuthorVkId:   201,
		UserAuthorName:   "Vasiliy Pupkin",
		UserAuthorAvatar: "https://yandex.ru/logo.png",
		UserExecutorVkId: pointy.Uint32(202),
		LocDep:           "Общежитие №10",
		LocArr:           "УЛК",
		DateTimeArr:      *dateTimeArr,
		Item:             "Зачётная книжка",
		MinPrice:         500,
		Comment:          "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, user_author_vk_id, user_author_name, user_author_avatar, user_executor_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad").
		WithArgs(expectedAd.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
				"user_author_avatar", "user_executor_vk_id", "loc_dep", "loc_dep", "date_time_arr",
				"item", "min_price", "comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
					expectedAd.UserAuthorAvatar, expectedAd.UserExecutorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Select(expectedAd.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Select_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	const id uint32 = 1

	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, user_author_vk_id, user_author_name, user_author_avatar, user_executor_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	resultAd, resultErr := adRepository.Select(id)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultAd)

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
	ad := &models.Ad{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:               ad.Id,
		UserAuthorId:     ad.UserAuthorId,
		UserAuthorVkId:   201,
		UserAuthorName:   "Vasiliy Pupkin",
		UserAuthorAvatar: "https://yandex.ru/logo.png",
		LocDep:           ad.LocDep,
		LocArr:           ad.LocArr,
		DateTimeArr:      ad.DateTimeArr,
		Item:             ad.Item,
		MinPrice:         ad.MinPrice,
		Comment:          ad.Comment,
	}

	sqlmock_.
		ExpectQuery("UPDATE ad").
		WithArgs(expectedAd.Id, expectedAd.LocDep, expectedAd.LocArr, time.Time(expectedAd.DateTimeArr),
			expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
				"user_author_avatar", "user_executor_vk_id", "loc_dep", "loc_dep", "date_time_arr", "item", "min_price",
				"comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
					expectedAd.UserAuthorAvatar, expectedAd.UserExecutorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Update(expectedAd)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Update_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:          1,
		LocDep:      "Общежитие №10",
		LocArr:      "УЛК",
		DateTimeArr: *dateTimeArr,
		Item:        "Зачётная книжка",
		MinPrice:    500,
		Comment:     "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("UPDATE ad").
		WithArgs(ad.Id, ad.LocDep, ad.LocArr, time.Time(ad.DateTimeArr), ad.Item, ad.MinPrice, ad.Comment).
		WillReturnError(sql.ErrNoRows)

	resultAd, resultErr := adRepository.Update(ad)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Delete(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	dateTimeArr, err := timestamps.NewDateTime("22.11.2021 16:55")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:               1,
		UserAuthorId:     101,
		UserAuthorVkId:   201,
		UserAuthorName:   "Vasiliy Pupkin",
		UserAuthorAvatar: "https://yandex.ru/logo.png",
		LocDep:           "Общежитие №10",
		LocArr:           "УЛК",
		DateTimeArr:      *dateTimeArr,
		Item:             "Зачётная книжка",
		MinPrice:         500,
		Comment:          "Поеду на велосипеде",
	}

	sqlmock_.
		ExpectQuery("DELETE FROM ad").
		WithArgs(expectedAd.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
				"user_author_avatar", "user_executor_vk_id", "loc_dep", "loc_dep", "date_time_arr", "item", "min_price",
				"comment"}).
				AddRow(expectedAd.Id, expectedAd.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
					expectedAd.UserAuthorAvatar, expectedAd.UserExecutorVkId, expectedAd.LocDep, expectedAd.LocArr,
					time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment))

	resultAd, resultErr := adRepository.Delete(expectedAd.Id)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAd, resultAd)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_Delete_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	const id uint32 = 1

	sqlmock_.
		ExpectQuery("DELETE FROM ad").
		WithArgs(id).
		WillReturnError(sql.ErrNoRows)

	resultAd, resultErr := adRepository.Delete(id)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultAd)

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
	adsSearch := HandoverTesting.NewAdsSearch(101, "Общежитие", "СК", *dateTimeArr1, 1000)
	expectedAds := &models.Ads{
		&models.Ad{
			Id:               1,
			UserAuthorId:     101,
			UserAuthorVkId:   201,
			UserAuthorName:   "Vasiliy Pupkin",
			UserAuthorAvatar: "https://yandex.ru/logo.png",
			LocDep:           "Общежитие №10",
			LocArr:           "УЛК",
			DateTimeArr:      *dateTimeArr1,
			Item:             "Тубус",
			MinPrice:         500,
			Comment:          "Поеду на коньках",
		},
		&models.Ad{
			Id:               2,
			UserAuthorId:     102,
			UserAuthorVkId:   202,
			UserAuthorName:   "Pupok Vasiliev",
			UserAuthorAvatar: "https://yandex.ru/logo2.png",
			LocDep:           "Общежитие №9",
			LocArr:           "СК",
			DateTimeArr:      *dateTimeArr2,
			Item:             "Спортивная форма",
			MinPrice:         600,
			Comment:          "Поеду на роликах :)",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
		"user_author_avatar", "user_executor_vk_id", "loc_dep", "loc_dep", "date_time_arr", "item", "min_price",
		"comment"})
	for _, expectedAd := range *expectedAds {
		rows.AddRow(expectedAd.Id, expectedAd.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
			expectedAd.UserAuthorAvatar, expectedAd.UserExecutorVkId, expectedAd.LocDep, expectedAd.LocArr,
			time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment)
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, user_author_vk_id, user_author_name, user_author_avatar, user_executor_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad").
		WithArgs(adsSearch.UserAuthorId, adsSearch.LocDep, adsSearch.LocArr, time.Time(*adsSearch.DateTimeArr),
			adsSearch.MaxPrice).
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
			Id:             1,
			UserAuthorId:   101,
			UserAuthorVkId: 201,
			LocDep:         "Общежитие №10",
			LocArr:         "СК",
			DateTimeArr:    *dateTimeArr1,
			MinPrice:       500,
			Comment:        "Поеду на коньках",
		},
		&models.Ad{
			Id:             2,
			UserAuthorId:   102,
			UserAuthorVkId: 202,
			LocDep:         "Общежитие №9",
			LocArr:         "СК",
			DateTimeArr:    *dateTimeArr2,
			MinPrice:       600,
			Comment:        "Поеду на роликах :)",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_author_id", "user_author_vk_id", "user_author_name",
		"user_author_avatar", "user_executor_vk_id", "loc_dep", "loc_dep", "date_time_arr", "item", "min_price",
		"comment"})
	for _, expectedAd := range *expectedAds {
		rows.AddRow(expectedAd.Id, expectedAd.UserAuthorId, expectedAd.UserAuthorVkId, expectedAd.UserAuthorName,
			expectedAd.UserAuthorAvatar, expectedAd.UserExecutorVkId, expectedAd.LocDep, expectedAd.LocArr,
			time.Time(expectedAd.DateTimeArr), expectedAd.Item, expectedAd.MinPrice, expectedAd.Comment)
	}
	sqlmock_.
		ExpectQuery("SELECT id, user_author_id, user_author_vk_id, user_author_name, user_author_avatar, user_executor_vk_id, loc_dep, loc_arr, date_time_arr, item, min_price, comment FROM ad").
		WillReturnRows(rows)

	resultAds, resultErr := adRepository.SelectArray(adsSearch)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAds, resultAds)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_InsertAdUserExecution(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	expectedAdUserExecution := &models.AdUserExecution{
		AdId:           1,
		UserExecutorId: 101,
	}

	sqlmock_.
		ExpectQuery("INSERT INTO ad_user_execution").
		WithArgs(expectedAdUserExecution.AdId, expectedAdUserExecution.UserExecutorId).
		WillReturnRows(
			sqlmock.NewRows([]string{"ad_id", "user_executor_id"}).
				AddRow(expectedAdUserExecution.AdId, expectedAdUserExecution.UserExecutorId))

	resultAdUserExecution, resultErr := adRepository.InsertAdUserExecution(expectedAdUserExecution)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAdUserExecution, resultAdUserExecution)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_SelectAdUserExecution(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	expectedAdUserExecution := &models.AdUserExecution{
		AdId:           1,
		UserExecutorId: 101,
	}

	sqlmock_.
		ExpectQuery("SELECT ad_id, user_executor_id FROM ad_user_execution").
		WithArgs(expectedAdUserExecution.AdId).
		WillReturnRows(
			sqlmock.NewRows([]string{"ad_id", "user_executor_id"}).
				AddRow(expectedAdUserExecution.AdId, expectedAdUserExecution.UserExecutorId))

	resultAdUserExecution, resultErr := adRepository.SelectAdUserExecution(expectedAdUserExecution.AdId)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAdUserExecution, resultAdUserExecution)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_SelectAdUserExecution_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	const adId uint32 = 1

	sqlmock_.
		ExpectQuery("SELECT ad_id, user_executor_id FROM ad_user_execution").
		WithArgs(adId).
		WillReturnError(sql.ErrNoRows)

	resultAdUserExecution, resultErr := adRepository.SelectAdUserExecution(adId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultAdUserExecution)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_DeleteAdUserExecution(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	expectedAdUserExecution := &models.AdUserExecution{
		AdId:           1,
		UserExecutorId: 101,
	}

	sqlmock_.
		ExpectQuery("DELETE FROM ad_user_execution").
		WithArgs(expectedAdUserExecution.AdId).
		WillReturnRows(
			sqlmock.NewRows([]string{"ad_id", "user_executor_id"}).
				AddRow(expectedAdUserExecution.AdId, expectedAdUserExecution.UserExecutorId))

	resultAdUserExecution, resultErr := adRepository.DeleteAdUserExecution(expectedAdUserExecution.AdId)
	assert.Nil(t, resultErr)
	assert.Equal(t, expectedAdUserExecution, resultAdUserExecution)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}

func TestAdRepository_DeleteAdUserExecution_notFound(t *testing.T) {
	db, sqlmock_, err := sqlmock.New()
	assert.Nil(t, err)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	adRepository := repository.NewAdRepositoryImpl(db)

	const adId uint32 = 1

	sqlmock_.
		ExpectQuery("DELETE FROM ad_user_execution").
		WithArgs(adId).
		WillReturnError(sql.ErrNoRows)

	resultAdUserExecution, resultErr := adRepository.DeleteAdUserExecution(adId)
	assert.Equal(t, resultErr, consts.RepErrNotFound)
	assert.Nil(t, resultAdUserExecution)

	assert.Nil(t, sqlmock_.ExpectationsWereMet())
}
