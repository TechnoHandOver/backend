package models

import (
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
)

type Ads struct { //TODO: нужны валидаторы моделей в handler'ах
	Id             uint32   `json:"id"`
	UserAuthorVkId uint32   `json:"userAuthorVkId"`
	LocDep         string   `json:"locDep"`
	LocArr         string   `json:"locArr"`
	DateTimeArr    DateTime `json:"dateTimeArr"`
	MinPrice       uint32   `json:"minPrice,omitempty"`
	Comment        string   `json:"comment,omitempty"`
}

type Adses []*Ads //TODO: Ad, Ads, не Adses... (переименовать)

type AdsUpdate struct {
	LocDep      string   `json:"locDep"`
	LocArr      string   `json:"locArr"`
	DateTimeArr DateTime `json:"dateTimeArr"`
	MinPrice    uint32   `json:"minPrice,omitempty"`
	Comment     string   `json:"comment,omitempty"`
}

type AdsSearch struct {
	LocDep      string   `query:"loc_dep"`
	LocArr      string   `query:"loc_arr"`
	DateTimeArr DateTime //`query:"date_time_arr"`
	MaxPrice    uint32   `query:"max_price"`
}
