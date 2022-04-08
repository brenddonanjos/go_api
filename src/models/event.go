package models

import "time"

type Event struct {
	Id            int       `json:"id" db:"id"`
	Provider      string    `json:"provider" db:"provider"`
	SpaceFlightId int       `json:"spaceFlightId" db:"spaceFlightId"`
	ArticleId     int       `json:"articleId" db:"articleId"`
	CreatedAt     time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updatedAt"`
}

type Events []Event

type SpaceFlightEvent struct {
	SpaceFlightId int    `json:"id" db:"id"`
	Provider      string `json:"provider" db:"provider"`
}

type SpaceFlightEvents []SpaceFlightEvent
