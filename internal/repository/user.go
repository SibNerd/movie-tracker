package repository

import (
	"movietracker/internal/db"
	entity "movietracker/internal/entities"

	"github.com/google/uuid"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(data entity.User) error {
	err := ur.db.CreateUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUser(id string) error {
	return nil
}

func (ur *UserRepository) GetUserShows(userId uuid.UUID) ([]entity.UserShow, error) {
	var shows []entity.UserShow

	shows, err := ur.db.GetUserShows(userId.String())
	if err != nil {
		return []entity.UserShow{}, err
	}

	return shows, nil
}

func (ur *UserRepository) GetUserShow(userShow entity.UserSingleShow) (entity.UserShow, error) {
	var shows entity.UserShow

	shows, err := ur.db.GetUserShow(userShow)
	if err != nil {
		return entity.UserShow{}, err
	}

	return shows, nil
}

func (ur *UserRepository) GetUserWatchlist(userId uuid.UUID) ([]entity.UserShow, error) {
	var shows []entity.UserShow

	shows, err := ur.db.GetUserShows(userId.String())
	if err != nil {
		return []entity.UserShow{}, err
	}

	return shows, nil
}

func (ur *UserRepository) AddShowToWatchlist(show entity.UserSingleShow) error {
	err := ur.db.AddShowToWatchlist(show)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) AddShowToList(show entity.UserShow) error {
	err := ur.db.AddShowToList(show)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) RateShow(userId, showId string, rating int) error {
	user, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	show, err := uuid.Parse(showId)
	if err != nil {
		return err
	}

	err = ur.db.RateShow(user, show, rating)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) MarkShowAsWatched(show entity.UserSingleShow) error {
	err := ur.db.MarkAsWatched(show)
	if err != nil {
		return err
	}

	return nil
}
