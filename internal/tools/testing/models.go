package testing

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
)

func NewAdsSearch(userAuthorId uint32, locDep, locArr string, dateTimeArr timestamps.DateTime, maxPrice uint32) *models.AdsSearch {
	return &models.AdsSearch{
		UserAuthorId: &userAuthorId,
		LocDep:       &locDep,
		LocArr:       &locArr,
		DateTimeArr:  &dateTimeArr,
		MaxPrice:     &maxPrice,
	}
}

func NewAdsSearchByUserAuthorId(userAuthorId uint32) *models.AdsSearch {
	return &models.AdsSearch{
		UserAuthorId: &userAuthorId,
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
