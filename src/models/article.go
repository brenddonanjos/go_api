package models

import "time"

type Article struct {
	Id            int       `json:"id" db:"id"`
	Featured      bool      `json:"featured" db:"featured"`
	Title         string    `json:"title" db:"title" validate:"required"`
	Url           string    `json:"url" db:"url" validate:"required"`
	ImageUrl      string    `json:"imageUrl" db:"imageUrl"`
	NewsSite      string    `json:"newsSite" db:"newsSite"`
	Summary       string    `json:"summary" db:"summary"`
	SpaceFlightId int       `json:"spaceFlightId" db:"spaceFlightId"`
	PublishedAt   time.Time `json:"publishedAt" db:"publishedAt"`
	CreatedAt     time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updatedAt"`
	Launches      Launches  `json:"launches"`
	Events        Events    `json:"events"`
}

type Articles []Article

type SpaceFlightArticle struct {
	SpaceFlightId int                 `json:"id" db:"id"`
	Featured      bool                `json:"featured" db:"featured"`
	Title         string              `json:"title" db:"title"`
	Url           string              `json:"url" db:"url"`
	ImageUrl      string              `json:"imageUrl" db:"imageUrl"`
	NewsSite      string              `json:"newsSite" db:"newsSite"`
	Summary       string              `json:"summary" db:"summary"`
	PublishedAt   string              `json:"publishedAt" db:"publishedAt"`
	UpdatedAt     string              `json:"updatedAt" db:"updatedAt"`
	Launches      SpaceFlightLaunches `json:"launches"`
	Events        SpaceFlightEvents   `json:"events"`
}

type SpaceFlightArticles []SpaceFlightArticle
