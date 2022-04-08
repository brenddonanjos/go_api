package webservices

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/brenddonanjos/go_api/src/models"
)

const (
	SFN_HOST = "https://api.spaceflightnewsapi.net/v3"
)

func ArticlesAll() (articles models.Articles, err error) {
	res, err := http.Get(SFN_HOST + "/articles") //Get all articles from SpaceFlightNews API
	if err != nil {
		return
	}
	defer res.Body.Close()

	sfArticles := models.SpaceFlightArticles{}
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(data), &sfArticles)
	if err != nil {
		return
	}

	//mount full models.Article struct (ready to save on db)
	for _, sfa := range sfArticles {
		dateLayout := "2006-01-02T15:04:05.000Z"
		publishedAt, _ := time.Parse(dateLayout, sfa.PublishedAt)
		updatedAt, _ := time.Parse(dateLayout, sfa.UpdatedAt)

		a := models.Article{}
		a.Featured = sfa.Featured
		a.Title = sfa.Title
		a.Url = sfa.Url
		a.ImageUrl = sfa.ImageUrl
		a.NewsSite = sfa.NewsSite
		a.Summary = sfa.Summary
		a.SpaceFlightId = sfa.SpaceFlightId
		a.PublishedAt = publishedAt
		a.UpdatedAt = updatedAt
		a.CreatedAt = time.Now()
		//launches
		a.Launches = []models.Launche{}
		if len(sfa.Launches) > 0 {
			ls := models.Launches{}
			for _, sfl := range sfa.Launches {
				l := models.Launche{}
				l.Provider = sfl.Provider
				l.SpaceFlightId = sfl.SpaceFlightId
				l.CreatedAt = time.Now()
				l.UpdatedAt = time.Now()
				ls = append(ls, l)
			}
			a.Launches = ls
		}
		//events
		a.Events = []models.Event{}
		if len(sfa.Events) > 0 {
			es := models.Events{}
			for _, sfl := range sfa.Events {
				e := models.Event{}
				e.Provider = sfl.Provider
				e.SpaceFlightId = sfl.SpaceFlightId
				e.CreatedAt = time.Now()
				e.UpdatedAt = time.Now()
				es = append(es, e)
			}
			a.Events = es
		}
		articles = append(articles, a)
	}

	return
}
