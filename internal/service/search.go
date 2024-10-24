package service

import (
	"fmt"
	entity "movietracker/internal/entities"

	"github.com/kataras/iris/v12"
)

type ResponseError struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

type SearchService interface {
	SearchShow(params entity.SearchParams) (entity.ShowsSearchResult, error)
}

type ServiceHandler struct {
	Service Service
}

func NewServiceHandler(e *iris.Application, svc *Service) {
	handler := &ServiceHandler{
		Service: *svc,
	}
	e.Post("/search", handler.SearchShow)
}

func (sh *ServiceHandler) SearchShow(ctx iris.Context) {
	var params entity.SearchParams

	err := ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	shows, err := sh.searchShow(params)
	if err != nil {
		fmt.Printf("error handling request")
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	ctx.JSON(iris.Map{"shows": shows})
}

func (sh *ServiceHandler) searchShow(params entity.SearchParams) (entity.ShowsSearchResult, error) {
	shows, err := sh.Service.searchRepo.SearchShow(params)
	if err != nil {
		return entity.ShowsSearchResult{}, err
	}

	return shows, nil
}
