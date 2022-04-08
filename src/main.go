package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brenddonanjos/go_api/src/actions"
	"github.com/brenddonanjos/go_api/src/database"
	"github.com/brenddonanjos/go_api/src/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
)

func main() {
	e := echo.New()

	//TODO: Code Middleware here
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //enable cors
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//DB instance
	db := database.Open()
	defer db.Close()

	//Validator
	e.Validator = &handlers.CustomValidator{
		Validator: validator.New(),
	}
	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Back-end Challenge 2021 🏅 - Space Flight News")
	})
	//Articles
	article := handlers.NewArticleHandler(db)
	e.GET("/articles", article.GetArticles)
	e.GET("/articles/:id", article.ShowArticles)
	e.POST("/articles", article.SetArticle)
	e.PUT("/articles/:id", article.UpdateArticle)
	e.DELETE("/articles/:id", article.DeleteArticle)
	e.GET("/articles/sync", article.SyncRoute)

	//create a sync cron
	cronExec(article)

	//Start aplication
	e.Logger.Fatal(e.Start(":8000"))
}

//cronExec executes a cron to sync spaceFlightNews API with db every day
func cronExec(a *handlers.Article) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	c := cron.New(cron.WithLocation(location))

	c.AddFunc("0 9 * * ?", func() { //start CRON at 0 min 9 hours every day
		err := a.SyncArticle()
		if err != nil {
			actions.ErrControl(err)
		}
	})

	c.Start()
	fmt.Println("Cron Started!")
}
