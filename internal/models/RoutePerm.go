package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type RoutePerm struct {
	Id             uint32    `json:"id"`
	UserAuthorVkId uint32    `json:"userAuthorVkId"`
	LocDep         string    `json:"locDep"`
	LocArr         string    `json:"locArr"`
	MinPrice       uint32    `json:"minPrice,omitempty"`
	EvenWeek       bool      `json:"evenWeek,omitempty"`
	OddWeek        bool      `json:"oddWeek,omitempty"`
	DayOfWeek      DayOfWeek `json:"dayOfWeek"`
	TimeDep        Time      `json:"timeDep"`
	TimeArr        Time      `json:"timeArr"`
}

type RoutesPerm []*RoutePerm
