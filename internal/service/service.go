package service

import (
	entity "movietracker/internal/entities"
)

type SearchRepository interface {
	SearchShow(params entity.SearchParams) (entity.ShowsSearchResult, error)
}

type Service struct {
	searchRepo SearchRepository
	//userRepo  UserRepository
}

func NewService(sr SearchRepository) *Service {
	return &Service{
		//userRepo: ur,
		searchRepo: sr,
	}
}
