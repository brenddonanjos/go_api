package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/brenddonanjos/go_api/src/actions"
	"github.com/brenddonanjos/go_api/src/webservices"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Article struct {
	ActionsArticle *actions.Article
	*sqlx.DB
}

// NewArticleHandler is a function to instance a new Article with db conn
func NewArticleHandler(db *sqlx.DB) *Article {
	return &Article{
		ActionsArticle: actions.NewArticleAction(db),
	}
}

func (a Article) GetArticles(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))

	if page == 0 {
		page = 1
	}

	articles, err := a.ActionsArticle.All(page)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, articles)
}

func (a Article) ShowArticles(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	article, err := a.ActionsArticle.Find(id)

	if err == sql.ErrNoRows {
		return c.String(http.StatusOK, "{}")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, article)
}

func (a Article) SetArticle(c echo.Context) error {
	//bind params on struct
	if err := c.Bind(&a.ActionsArticle.Article); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Error(err))
	}
	a.ActionsArticle.Article.PublishedAt = time.Now()
	a.ActionsArticle.Article.CreatedAt = time.Now()
	a.ActionsArticle.Article.UpdatedAt = time.Now()

	//validate required fields
	if err := c.Validate(a.ActionsArticle.Article); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, Error(err))
	}

	//creates new article
	res, err := a.ActionsArticle.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	insertId, _ := res.LastInsertId()
	a.ActionsArticle.Article.Id = int(insertId)
	return c.JSON(http.StatusCreated, Success("Article created successfuly", a.ActionsArticle.Article))
}

func (a Article) UpdateArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&a.ActionsArticle.Article); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Error(err))
	}
	a.ActionsArticle.Article.Id = id
	err := a.ActionsArticle.Update(id)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, Success("Article updated", a.ActionsArticle.Article))
}

func (a Article) DeleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.ActionsArticle.Delete(id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, "Article deleted successfuly!")
}

func (a Article) SyncRoute(c echo.Context) error {
	err := a.SyncArticle()

	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError)
	}
	c.String(http.StatusCreated, "sync finished!")

	return err
}

func (a Article) SyncArticle() (err error) {
	//Get results from spaceflight and convert on model article struct
	articles, err := webservices.ArticlesAll()

	//save on db using golang multi-thread processes (Goroutines)
	a.ActionsArticle.CheckAndSave(articles) //this function returns a channel if we wants manipulated it

	/*
		//save on db using on synchronous way because some databes providers only supports 1 connection/operation at time
		a.ActionsArticle.CheckAndSaveSynchronous(articles)
	*/

	return
}
