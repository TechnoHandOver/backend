package ad

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type Usecase interface {
	Create(ad_ *models.Ad) *response.Response
	Get(id uint32) *response.Response
	Update(ad_ *models.Ad) *response.Response
	Delete(userVkId uint32, id uint32) *response.Response
	Search(adsSearch *models.AdsSearch) *response.Response
}
