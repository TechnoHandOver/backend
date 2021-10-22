package models

import (
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
)

type Ads struct {
	Id             uint32   `json:"id"`
	UserAuthorVkId uint32   `json:"userAuthorVkId"`
	LocDep         string   `json:"locDep"`
	LocArr         string   `json:"locArr"`
	DateTimeArr    DateTime `json:"dateTimeArr"`
	MinPrice       uint32   `json:"minPrice,omitempty"`
	Comment        string   `json:"comment,omitempty"`
}

type Adses []*Ads

type AdsUpdate struct {
	LocDep         string   `json:"locDep"`
	LocArr         string   `json:"locArr"`
	DateTimeArr    DateTime `json:"dateTimeArr"`
	MinPrice       uint32   `json:"minPrice,omitempty"`
	Comment        string   `json:"comment,omitempty"`
}
