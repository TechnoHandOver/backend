package ads

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type AdsUsecase interface {
	Create(ads *models.Ads) *response.Response
	Get(id uint32) *response.Response
	Update(id uint32, adsUpdate *models.AdsUpdate) *response.Response
	List() *response.Response
	Search(adsSearch *models.AdsSearch) *response.Response
}

type AdsRepository interface {
	Insert(ads *models.Ads) (*models.Ads, error)
	Select(id uint32) (*models.Ads, error)
	Update(id uint32, adsUpdate *models.AdsUpdate) (*models.Ads, error)
	SelectArray() (*models.Adses, error)
	SelectArray2(adsSearch *models.AdsSearch) (*models.Adses, error)
}
