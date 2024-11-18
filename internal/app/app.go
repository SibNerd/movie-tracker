package app

import (
	"fmt"
	"movietracker/internal/config"
	"movietracker/internal/db/cockroach"
	"movietracker/internal/repository"
	"movietracker/internal/service"

	"github.com/caarlos0/env/v11"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type App struct {
	app    *iris.Application
	svc    *service.Service
	config config.Config
	logger *logrus.Logger
}

func NewApp() App {
	logger := logrus.New()

	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		newErr := fmt.Errorf("failed to parse config file: %e", err)
		logger.Panic(newErr)
	}
	app := iris.New()

	db, err := cockroach.ConnectDB(cfg.DatabaseDSN)
	if err != nil {
		newErr := fmt.Errorf("failed to connect to DB: %e", err)
		logger.Panic(newErr)
	}

	searcher := repository.NewSearchRepository(cfg.SearcherURL)
	user := repository.NewUserRepository(db)
	svc := service.NewService(searcher, user)
	service.NewServiceHandler(app, svc)

	return App{
		app:    app,
		svc:    svc,
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() {
	url := a.config.Host + ":" + a.config.Port
	a.logger.Fatal(a.app.Listen(url))
}
