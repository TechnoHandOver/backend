package models

import "time"

type Ads struct {
	Id           uint32    `json:"id,omitempty"`
	UserAuthorId uint32    `json:"userAuthorId,omitempty"`
	LocationFrom string    `json:"locationFrom"`
	LocationTo   string    `json:"locationTo"`
	TimeFrom     time.Time `json:"timeFrom"`
	TimeTo       time.Time `json:"timeTo"`
	MinPrice     uint32    `json:"minPrice,omitempty"`
	Comment      string    `json:"comment,omitempty"`
}

type Adses []*Ads
