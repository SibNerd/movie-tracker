package service

import (
	entity "movietracker/internal/entities"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type SearchRepository interface {
	SearchShow(params entity.SearchParams) (entity.ShowsSearchResult, error)
}

type UserRepository interface {
	CreateUser(data entity.User) error
	GetUser(id string) error

	GetUserShows(userId uuid.UUID) ([]entity.UserShow, error)
	GetUserShow(userShow entity.UserSingleShow) (entity.UserShow, error)
	AddShowToList(show entity.UserShow) error
	GetUserWatchlist(userId uuid.UUID) ([]entity.UserShow, error)
	AddShowToWatchlist(show entity.UserSingleShow) error
	RateShow(userId, showId string, rating int) error
	MarkShowAsWatched(show entity.UserSingleShow) error
}

type Service struct {
	searchRepo SearchRepository
	userRepo   UserRepository
}

func NewService(sr SearchRepository, ur UserRepository) *Service {
	return &Service{
		userRepo:   ur,
		searchRepo: sr,
	}
}

type ServiceHandler struct {
	Service Service
}

func NewServiceHandler(e *iris.Application, svc *Service) {
	handler := &ServiceHandler{
		Service: *svc,
	}

	e.Post("/search", handler.SearchShow)

	//e.Post("/save-show", handler.SaveShow)

	users := e.Party("/users")
	users.Post("/signup", handler.UserSignUp)
	users.Get("/login", handler.UserLogin)

	usershows := users.Party("/shows")
	usershows.Get("/all", handler.GetUserShows)
	usershows.Get("/watchlist", handler.GetUserWatchlist)
	usershows.Get("/show", handler.GetUserShow)
	usershows.Post("/watchlist", handler.AddToUserWatchlist)
	usershows.Put("/watchlist", handler.MarkAsWatched)
	usershows.Put("/rate", handler.RateShow)
	usershows.Post("/watched", handler.AddToWatched)
}
