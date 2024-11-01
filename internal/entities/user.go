package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

type UserShow struct {
	UserID      uuid.UUID `json:"userID"`
	ShowID      uuid.UUID `json:"showID"`
	IsWatched   bool      `json:"isWatched"`
	OnWatchlist bool      `json:"onWatchlist"`
	Rating      int       `json:"rating"`
	Note        string    `json:"note"`
}

type UserSingleShow struct {
	UserID uuid.UUID `json:"userID"`
	ShowID uuid.UUID `json:"showID"`
}
