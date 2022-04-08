package models

import "time"

type Launche struct {
	Id            int       `json:"id" db:"id"`
	Provider      string    `json:"provider" db:"provider"`
	SpaceFlightId string    `json:"spaceFlightId" db:"spaceFlightId"`
	ArticleId     int       `json:"articleId" db:"articleId"`
	CreatedAt     time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updatedAt"`
}

type Launches []Launche

type SpaceFlightLaunche struct {
	SpaceFlightId string `json:"id" db:"id"`
	Provider      string `json:"provider" db:"provider"`
}

type SpaceFlightLaunches []SpaceFlightLaunche
