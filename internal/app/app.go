package app

import (
	"movietracker/internal/config"
	"movietracker/internal/repository"
	"movietracker/internal/service"

	"github.com/caarlos0/env/v11"
	"github.com/kataras/iris/v12"
)

type App struct {
	app *iris.Application
	svc *service.Service
}

func NewApp() App {
	//logger := log.New()

	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		panic("failed to parse config file")
		//logger.Panic(err)
	}
	app := iris.New()

	searcher := repository.NewSearchRepository(cfg.SearcherURL)
	svc := service.NewService(searcher)
	service.NewServiceHandler(app, svc)

	return App{
		app: app,
		svc: svc,
	}
}

func (a *App) Run() {
	a.app.Listen(":8080")
}
