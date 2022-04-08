package actions

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/go_api/src/models"
	"github.com/jmoiron/sqlx"
)

type Event struct {
	Event models.Event
	Db    *sqlx.DB
}

func NewEventAction(db *sqlx.DB) *Event {
	return &Event{
		Db: db,
	}
}

func (e Event) NewEvent() (err error) {
	stmt, _ := e.Db.Preparex("INSERT INTO events (provider, articleId, spaceFlightId, createdAt, updatedAt) values (?, ?, ?, ?, ?)")
	_, err = stmt.Exec(e.Event.Provider, e.Event.ArticleId, e.Event.SpaceFlightId, time.Now(), time.Now())

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (e Event) FindByArticleId(id int) (events models.Events, err error) {
	stmt, _ := e.Db.Preparex("SELECT * FROM events WHERE articleId = ?")
	err = stmt.Select(&events, id)
	if err != nil {
		return
	}
	if events == nil {
		events = models.Events{}
	}
	return
}

func (e Event) CountEventSync(id int) (count int, err error) {
	err = e.Db.QueryRowx("SELECT COUNT(*) FROM events WHERE spaceFlightId = ?", id).Scan(&count)

	return
}
