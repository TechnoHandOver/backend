package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type Ad struct {
	Id             uint32   `json:"id"`
	UserAuthorVkId uint32   `json:"userAuthorVkId"`
	LocDep         string   `json:"locDep"`
	LocArr         string   `json:"locArr"`
	DateTimeArr    DateTime `json:"dateTimeArr"`
	Item           string   `json:"item"`
	MinPrice       uint32   `json:"minPrice,omitempty"`
	Comment        string   `json:"comment,omitempty"`
}

type Ads []*Ad

type AdsSearch struct {
	LocDep      string   `query:"loc_dep"`
	LocArr      string   `query:"loc_arr"`
	DateTimeArr DateTime `query:"date_time_arr"`
	MaxPrice    uint32   `query:"max_price"`
}
