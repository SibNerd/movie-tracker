package db

import (
	entity "movietracker/internal/entities"

	"github.com/google/uuid"
)

type DB interface {
	UserDB
	ShowDB
}

type UserDB interface {
	CreateUser(data entity.User) error
	GetUser(id string) (entity.User, error)
}

type ShowDB interface {
	GetUserShows(id string) ([]entity.UserShow, error)
	GetUserShow(us entity.UserSingleShow) (entity.UserShow, error)
	GetUserWatchlist(id string) ([]entity.UserShow, error)
	AddShowToWatchlist(show entity.UserSingleShow) error
	AddShowToList(show entity.UserShow) error
	RateShow(userId, showId uuid.UUID, rating int) error
	MarkAsWatched(show entity.UserSingleShow) error
}
