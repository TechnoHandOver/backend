package models

import (
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
	"reflect"
)

type AdsSearch struct {
	UserAuthorId    *uint32
	NotUserAuthorId *uint32
	LocDep          *string
	LocArr          *string
	DateTimeArr     *DateTime
	MaxPrice        *uint32
	Order           *AdsSearchOrder
}

type AdsSearchOrder int

const (
	AdsSearchOrderDateTimeArrDesc AdsSearchOrder = iota
	AdsSearchOrderDateTimeArrAsc
	AdsSearchOrderMinPriceDesc
	AdsSearchOrderMinPriceAsc
)

func ValidateAdsSearchOrder(field reflect.Value) interface{} {
	if order, ok := field.Interface().(AdsSearchOrder); ok {
		switch order {
		case AdsSearchOrderDateTimeArrDesc, AdsSearchOrderDateTimeArrAsc, AdsSearchOrderMinPriceDesc,
			AdsSearchOrderMinPriceAsc:
			return true
		}
	}
	return nil
}
