package cockroach

import (
	entity "movietracker/internal/entities"

	"github.com/google/uuid"
)

// User Table functions

func (db *Database) CreateUser(data entity.User) error {
	res := db.DB.
		Table("users").
		Create(&data)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *Database) GetUser(id string) (entity.User, error) {
	var user entity.User
	res := db.DB.
		Table("users").
		Where(entity.User{ID: uuid.MustParse(id)}).
		First(&user)
	if res.Error != nil {
		return entity.User{}, res.Error
	}

	return user, nil
}

// UserShow table functions

func (db *Database) GetUserShows(id string) ([]entity.UserShow, error) {
	var shows []entity.UserShow
	res := db.DB.
		Table("userShows").
		Where(entity.User{ID: uuid.MustParse(id)}).
		Find(&shows).
		Order("showId ASC")
	if res.Error != nil {
		return []entity.UserShow{}, res.Error
	}

	return shows, nil
}

func (db *Database) GetUserShow(us entity.UserSingleShow) (entity.UserShow, error) {
	var show entity.UserShow
	res := db.DB.
		Table("userShows").
		Where(entity.UserShow{UserID: us.UserID}).
		Where(entity.UserShow{ShowID: us.ShowID}).
		First(&show)
	if res.Error != nil {
		return entity.UserShow{}, res.Error
	}
	return show, nil
}

func (db *Database) GetUserWatchlist(id string) ([]entity.UserShow, error) {
	var shows []entity.UserShow
	res := db.DB.
		Table("userShows").
		Where(entity.User{ID: uuid.MustParse(id)}).
		Where(entity.UserShow{OnWatchlist: true}).
		Find(&shows).
		Order("showId ASC")
	if res.Error != nil {
		return []entity.UserShow{}, res.Error
	}

	return shows, nil
}

func (db *Database) AddShowToWatchlist(show entity.UserSingleShow) error {
	res := db.DB.
		Table("userShows").
		Where(&entity.UserShow{
			UserID: show.UserID,
			ShowID: show.ShowID}).
		Update("onWatchlist", true)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *Database) AddShowToList(show entity.UserShow) error {
	res := db.DB.
		Table("userShows").
		Create(&show)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *Database) RateShow(userId, showId uuid.UUID, rating int) error {
	res := db.DB.
		Table("userShows").
		Where(&entity.UserShow{
			UserID: userId,
			ShowID: showId}).
		Update("rating", rating)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *Database) MarkAsWatched(show entity.UserSingleShow) error {
	res := db.DB.
		Table("userShows").
		Where(&entity.UserShow{UserID: show.UserID, ShowID: show.ShowID}).
		Update("isWatched", true)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
