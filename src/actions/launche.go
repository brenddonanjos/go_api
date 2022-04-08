package actions

import (
	"fmt"
	"time"

	"github.com/brenddonanjos/go_api/src/models"
	"github.com/jmoiron/sqlx"
)

type Launche struct {
	Launche models.Launche
	Db      *sqlx.DB
}

func NewLauncheAction(db *sqlx.DB) *Launche {
	return &Launche{
		Db: db,
	}
}

func (l Launche) NewLaunche() (err error) {
	stmt, _ := l.Db.Preparex("INSERT INTO launches (provider, articleId, spaceFlightId, createdAt, updatedAt) values (?, ?, ?, ?, ?)")
	_, err = stmt.Exec(l.Launche.Provider, l.Launche.ArticleId, l.Launche.SpaceFlightId, time.Now(), time.Now())

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (l Launche) FindByArticleId(id int) (launches models.Launches, err error) {
	stmt, _ := l.Db.Preparex("SELECT * FROM launches WHERE articleId = ?")
	err = stmt.Select(&launches, id)
	if err != nil {
		return
	}
	if launches == nil {
		launches = models.Launches{}
	}
	return
}

func (l Launche) CountLaunchSync(id int) (count int, err error) {
	err = l.Db.QueryRowx("SELECT COUNT(*) FROM launches WHERE spaceFlightId = ?", id).Scan(&count)

	return
}
