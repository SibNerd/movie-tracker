package service

import (
	"fmt"
	entity "movietracker/internal/entities"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func (sh *ServiceHandler) UserSignUp(ctx iris.Context) {
	var user entity.User

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	user.ID = uuid.New()
	err = sh.Service.userRepo.CreateUser(user)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) UserLogin(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) GetUserShows(ctx iris.Context) {
	var user entity.User

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	show, err := sh.Service.userRepo.GetUserShows(user.ID)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	ctx.JSON(iris.Map{"show": show})
}

func (sh *ServiceHandler) GetUserShow(ctx iris.Context) {
	var req entity.UserSingleShow

	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	show, err := sh.Service.userRepo.GetUserShow(req)
	if err != nil {
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Error: %e", err)})
	}

	ctx.JSON(iris.Map{"show": show})
}

func (sh *ServiceHandler) AddToWatched(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) GetUserWatchlist(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) AddToUserWatchlist(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) MarkAsWatched(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}

func (sh *ServiceHandler) RateShow(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "user signed up"})
}
