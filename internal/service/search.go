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

func (sh *ServiceHandler) SearchShow(ctx iris.Context) {
	var params entity.SearchParams

	err := ctx.ReadJSON(&params)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	shows, err := sh.Service.searchRepo.SearchShow(params)
	if err != nil {
		fmt.Printf("error handling request")
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	ctx.JSON(iris.Map{"shows": shows})
}
