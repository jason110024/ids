package dtos

import (
	"net/http"
	"time"

	"github.com/adh-partnership/ids/backend/internal/domain/airports"
)

type AirportPatch struct {
	ATIS             *string    `json:"atis"`
	ATISTime         *time.Time `json:"atis_time"`
	ArrivalATIS      *string    `json:"arrival_atis"`
	ArrivalATISTime  *time.Time `json:"arrival_atis_time"`
	DepartureRunways *string    `json:"departure_runways"`
	ArrivalRunways   *string    `json:"arrival_runways"`
}

func (a *AirportPatch) Bind(r *http.Request) error {
	return nil
}

func (a *AirportPatch) MergeInto(airport *airports.Airport) {
	if a.ATIS != nil {
		airport.ATIS = *a.ATIS
	}
	if a.ATISTime != nil {
		airport.ATISTime = a.ATISTime
	}
	if a.ArrivalATIS != nil {
		airport.ArrivalATIS = *a.ArrivalATIS
	}
	if a.ArrivalATISTime != nil {
		airport.ArrivalATISTime = a.ArrivalATISTime
	}
	if a.DepartureRunways != nil {
		airport.DepartureRunways = *a.DepartureRunways
	}
	if a.ArrivalRunways != nil {
		airport.ArrivalRunways = *a.ArrivalRunways
	}
}
