package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type RoutePerm struct {
	Id           uint32    `json:"id"`
	UserAuthorId uint32    `json:"-"`
	LocDep       string    `json:"locDep"`
	LocArr       string    `json:"locArr"`
	MinPrice     uint32    `json:"minPrice"`
	EvenWeek     bool      `json:"evenWeek"`
	OddWeek      bool      `json:"oddWeek"`
	DayOfWeek    DayOfWeek `json:"dayOfWeek"`
	TimeDep      Time      `json:"timeDep"`
	TimeArr      Time      `json:"timeArr"`
}

type RoutesPerm []*RoutePerm
