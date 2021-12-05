package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type RouteTmp struct {
	Id           uint32   `json:"id"`
	UserAuthorId uint32   `json:"-"`
	LocDep       string   `json:"locDep"`
	LocArr       string   `json:"locArr"`
	MinPrice     uint32   `json:"minPrice"`
	DateTimeDep  DateTime `json:"dateTimeDep"`
	DateTimeArr  DateTime `json:"dateTimeArr"`
}

type RoutesTmp []*RouteTmp
