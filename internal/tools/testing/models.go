package testing

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
)

func NewAdsSearch(userAuthorVkId uint32, locDep, locArr string, dateTimeArr timestamps.DateTime, maxPrice uint32) *models.AdsSearch {
	return &models.AdsSearch{
		UserAuthorVkId: &userAuthorVkId,
		LocDep:         &locDep,
		LocArr:         &locArr,
		DateTimeArr:    &dateTimeArr,
		MaxPrice:       &maxPrice,
	}
}

func NewAdsSearchByUserAuthorVkId(userAuthorVkId uint32) *models.AdsSearch {
	return &models.AdsSearch{
		UserAuthorVkId: &userAuthorVkId,
	}
}

func NewAdsSearchBySecondaryFields(locDep, locArr string, dateTimeArr timestamps.DateTime, maxPrice uint32) *models.AdsSearch {
	return &models.AdsSearch{
		LocDep:      &locDep,
		LocArr:      &locArr,
		DateTimeArr: &dateTimeArr,
		MaxPrice:    &maxPrice,
	}
}
