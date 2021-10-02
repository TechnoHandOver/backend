package models

import "time"

type Ads struct {
	Id           uint32    `json:"id"`
	UserAuthorId uint32    `json:"userAuthorId"`
	LocationFrom string    `json:"locationFrom"`
	LocationTo   string    `json:"locationTo"`
	TimeFrom     time.Time `json:"timeFrom"`
	TimeTo       time.Time `json:"timeTo"`
	MinPrice     uint32    `json:"minPrice"`
	Comment      string    `json:"comment"`
}

type Adses []*Ads
