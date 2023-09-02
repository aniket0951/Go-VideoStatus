package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"reflect"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	CreateAdminUser(*gin.Context)
	LoginAdminUser(*gin.Context)
	UpdateUserAccountStatus(*gin.Context)
	FetchAllUsers(*gin.Context)
}

type userController struct {
	userServ services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		userServ: service,
	}
}

func (cont *userController) CreateAdminUser(ctx *gin.Context) {
	var req dto.CreateAdminRequestParams

	if err := ctx.ShouldBindJSON(&req); helper.CheckError(err, ctx) {
		return
	}

	new_user, err := cont.userServ.CreateAdminUser(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.USER_REGISTRATION_SUCCESS, new_user, helper.USER_DATA)
	ctx.JSON(http.StatusCreated, response)
}

func (cont *userController) LoginAdminUser(ctx *gin.Context) {
	var req dto.LoginUserRequestParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	user, err := cont.userServ.LoginAdminUser(req)

	if err != nil {
		if err == sql.ErrNoRows {
			response := helper.BuildFailedResponse(helper.FETCHED_FAILED, errors.New("user not found").Error(), helper.EmptyObj{}, helper.DATA)
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		if helper.CheckError(err, ctx) {
			return
		}
	}

	response := helper.BuildSuccessResponse(helper.LOGIN_SUCCESS, user, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (cont *userController) UpdateUserAccountStatus(ctx *gin.Context) {
	var req dto.UpdateUserAccountStatusRequestParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if reflect.TypeOf(*req.AccountStatus).Kind() != reflect.Bool {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, errors.New("account status not found").Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// check user id is Valid or not
	user_id, err := uuid.Parse(req.Id)

	if helper.CheckError(err, ctx) {
		return
	}

	err = cont.userServ.UpdateUserAccountStatus(user_id, *req.AccountStatus)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("user account status has been changed", helper.EmptyObj{}, helper.DATA)
	ctx.JSON(http.StatusAccepted, response)
}

func (cont *userController) FetchAllUsers(ctx *gin.Context) {
	var req dto.GetUsersRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	users, err := cont.userServ.FetchAllUsers(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, users, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)
}
