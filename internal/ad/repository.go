package ad

import "github.com/TechnoHandOver/backend/internal/models"

type Repository interface {
	Insert(ad_ *models.Ad) (*models.Ad, error)
	Select(id uint32) (*models.Ad, error)
	Update(ad_ *models.Ad) (*models.Ad, error)
	SelectArray(adsSearch *models.AdsSearch) (*models.Ads, error)
	Delete(id uint32) (*models.Ad, error)
	InsertAdUserExecution(adUserExecution *models.AdUserExecution) (*models.AdUserExecution, error)
	SelectAdUserExecution(adId uint32) (*models.AdUserExecution, error)
	DeleteAdUserExecution(adId uint32) (*models.AdUserExecution, error)
}
