package models

import . "github.com/TechnoHandOver/backend/internal/models/timestamps"

type AdsSearch struct {
	UserAuthorId *uint32
	LocDep       *string
	LocArr       *string
	DateTimeArr  *DateTime
	MaxPrice     *uint32
}
