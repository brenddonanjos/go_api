package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brenddonanjos/go_api/src/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	articleJSON = `{"featured":false,"title":"teste1","url":"teste1","imageUrl":"teste1","newsSite":"teste1","summary":"teste1"}`
)

func TestSetArticle(t *testing.T) {
	db := database.Open()
	defer db.Close()
	e := echo.New()
	//Validator
	e.Validator = &CustomValidator{
		Validator: validator.New(),
	}
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	a := NewArticleHandler(db)
	if assert.NoError(t, a.SetArticle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUpdateArticle(t *testing.T) {
	db := database.Open()
	defer db.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/articles/", strings.NewReader(articleJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	a := NewArticleHandler(db)

	if assert.NoError(t, a.UpdateArticle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestShowArticle(t *testing.T) {
	db := database.Open()
	defer db.Close()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/articles/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	a := NewArticleHandler(db)
	// Assertions
	if assert.NoError(t, a.ShowArticles(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteArticle(t *testing.T) {
	db := database.Open()
	defer db.Close()

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/articles/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	a := NewArticleHandler(db)
	// Assertions
	if assert.NoError(t, a.DeleteArticle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Article deleted successfuly!", rec.Body.String())
	}
}
