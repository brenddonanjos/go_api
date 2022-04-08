package actions

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/brenddonanjos/go_api/src/models"
	"github.com/jmoiron/sqlx"
)

//Articles struct repository
type Article struct {
	Article models.Article
	Db      *sqlx.DB
}

func NewArticleAction(db *sqlx.DB) *Article {
	return &Article{
		Db: db,
	}
}

func (a Article) All(page int) (articles models.Articles, err error) {
	perPage := 10
	sql := "SELECT * FROM articles ORDER BY id DESC"

	offset := (page - 1) * perPage
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, offset)

	rows, err := a.Db.Queryx(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	articles = []models.Article{}
	for rows.Next() {
		var article models.Article
		err = rows.StructScan(&article)
		if err != nil {
			return
		}
		a.relationships(&article)

		articles = append(articles, article)
	}

	return
}

func (a Article) Find(id int) (article models.Article, err error) {
	row := a.Db.QueryRowx("SELECT * FROM articles WHERE id = ? LIMIT 1", id)
	err = row.StructScan(&article)
	a.relationships(&article)
	return
}

func (a Article) New() (result sql.Result, err error) {
	stmt, _ := a.Db.Preparex("INSERT INTO articles (featured, title, url, imageUrl, newsSite, summary, spaceFlightId, publishedAt, createdAt, updatedAt) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	res, err := stmt.Exec(a.Article.Featured, a.Article.Title, a.Article.Url, a.Article.ImageUrl, a.Article.NewsSite, a.Article.Summary, a.Article.SpaceFlightId, a.Article.PublishedAt, a.Article.CreatedAt, a.Article.UpdatedAt)

	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func (a Article) Update(id int) (err error) {
	stmt, _ := a.Db.Preparex("UPDATE articles SET featured = ?, title = ?, url = ?, imageUrl = ?, newsSite = ?, summary = ? WHERE id = ?")
	_, err = stmt.Exec(a.Article.Featured, a.Article.Title, a.Article.Url, a.Article.ImageUrl, a.Article.NewsSite, a.Article.Summary, id)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (a Article) Delete(id int) (err error) {
	stmt, _ := a.Db.Preparex("DELETE FROM articles WHERE id = ?")
	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err)
	}
	return
}

func (a Article) CountBySpaceFlightId(id int) (count int, err error) {
	err = a.Db.QueryRowx("SELECT COUNT(*) FROM articles WHERE spaceFlightId = ?", id).Scan(&count)

	return
}

func (a Article) relationships(article *models.Article) {
	//get launches
	l := NewLauncheAction(a.Db)
	launches, _ := l.FindByArticleId(article.Id)
	article.Launches = launches
	//get events
	ev := NewEventAction(a.Db)
	events, _ := ev.FindByArticleId(article.Id)
	article.Events = events
}

//CheckAndSave Execute saves on db in a multi thread way (using goroutines)
func (article Article) CheckAndSave(articles models.Articles) chan models.Article {

	channel := make(chan models.Article)

	for _, a := range articles {
		article.Article = a
		go func(art *Article) { //start goroutine

			count, err := art.CountBySpaceFlightId(art.Article.SpaceFlightId) //check if exists on db
			if err != nil {
				ErrControl(err)
			}

			if count == 0 {
				res, error := art.New() //save if not exists
				if error != nil {
					ErrControl(error)
				}
				articleid, _ := res.LastInsertId()
				art.Article.Id = int(articleid)
				channel <- art.Article //store on channel

				//events
				if len(art.Article.Events) > 0 {
					for _, e := range art.Article.Events {
						e.ArticleId = int(articleid)
						var event = NewEventAction(art.Db)
						event.Event = e
						err = event.NewEvent()
						if err != nil {
							ErrControl(err)
						}
					}
				}
				//launches
				if len(art.Article.Launches) > 0 {
					for _, l := range art.Article.Launches {
						l.ArticleId = int(articleid)
						var launche = Launche{
							Launche: l,
							Db:      art.Db,
						}
						err = launche.NewLaunche()
						if err != nil {
							ErrControl(err)
						}
					}
				}

			}
		}(&article)
	}
	return channel
}

//CheckAndSaveSynchronous Execute saves on db in a Synchronous way
func (article Article) CheckAndSaveSynchronous(articles models.Articles) models.Articles {
	storedArticles := models.Articles{}
	for _, a := range articles {
		article.Article = a

		count, err := article.CountBySpaceFlightId(article.Article.SpaceFlightId) //check if exists on db
		if err != nil {
			ErrControl(errors.New("Article sync failed: " + err.Error()))
		}

		if count == 0 {
			res, error := article.New() //save if not exists
			if error != nil {
				ErrControl(errors.New("Article sync failed: " + error.Error()))
			}
			articleid, _ := res.LastInsertId()
			article.Article.Id = int(articleid)

			storedArticles = append(storedArticles, a)

			//events
			if len(article.Article.Events) > 0 {
				for _, e := range article.Article.Events {
					e.ArticleId = int(articleid)
					var event = NewEventAction(article.Db)
					event.Event = e
					err = event.NewEvent()
					if err != nil {
						ErrControl(errors.New("Article sync failed: " + err.Error()))
					}
				}
			}
			//launches
			if len(article.Article.Launches) > 0 {
				for _, l := range article.Article.Launches {
					l.ArticleId = int(articleid)
					var launche = Launche{
						Launche: l,
						Db:      article.Db,
					}
					err = launche.NewLaunche()
					if err != nil {
						ErrControl(errors.New("Article sync failed: " + err.Error()))
					}
				}
			}

		}
	}
	return storedArticles
}
