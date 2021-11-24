package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type Ad struct {
	Id             uint32   `json:"id"`
	UserAuthorVkId uint32   `json:"userAuthorVkId"`
	LocDep         string   `json:"locDep"`
	LocArr         string   `json:"locArr"`
	DateTimeArr    DateTime `json:"dateTimeArr"`
	Item           string   `json:"item"`
	MinPrice       uint32   `json:"minPrice"`
	Comment        string   `json:"comment"`
}

type Ads []*Ad
