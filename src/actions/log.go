package actions

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/go_api/src/database"
)

type Log struct {
	Id        int       `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
}

func (l Log) NewLog() (err error) {
	db := database.Open()
	defer db.Close()

	stmt, _ := db.Preparex("INSERT INTO logs (message, type, createdAt, updatedAt) values (?, ?, ?, ?)")
	_, err = stmt.Exec(l.Message, l.Type, l.CreatedAt, l.UpdatedAt)

	if err != nil {
		fmt.Println(err)
	}

	return
}
