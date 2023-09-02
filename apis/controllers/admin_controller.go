package controllers

import (
	"errors"
	"net/http"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminController interface {
	UploadVideoByAdmin(ctx *gin.Context)
	GetVideoByAdmin(ctx *gin.Context)
	UpdateVideoStatus(ctx *gin.Context)
	FetchVerifyVideos(ctx *gin.Context)
}

type adminController struct {
	adminServ services.AdminService
}

func NewAdminController(service services.AdminService) AdminController {
	return &adminController{
		adminServ: service,
	}
}

func (adminCon *adminController) UploadVideoByAdmin(ctx *gin.Context) {
	file, _, _ := ctx.Request.FormFile("video_file")
	title := ctx.PostForm("title")
	uploaded_by := ctx.PostForm("uploaded_by")
	status := ctx.PostForm("status")

	if title == "" || status == "" {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	admin_user, err := uuid.Parse(uploaded_by)

	if helper.CheckError(err, ctx) {
		return
	}

	req := dto.UploadVideoByAdminRequestParam{
		Title:      title,
		UploadedBy: admin_user,
		Status:     status,
	}

	video, err := adminCon.adminServ.UploadVideoByAdmin(req, file)

	if helper.CheckError(err, ctx) {
		return
	}

	filePath := "http://localhost:8080/static/" + video.FileAddress
	video.FileAddress = filePath

	response := helper.BuildSuccessResponse("video has been uploaded", video, helper.DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) GetVideoByAdmin(ctx *gin.Context) {
	var req dto.GetVideoByAdminRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	if helper.TOKEN_ID == "" {
		response := helper.BuildFailedResponse(helper.FETCHED_FAILED, errors.New("unable to locate admin").Error(),
			helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	req.UserId = helper.TOKEN_ID

	videos, err := adminCon.adminServ.GetVideoByAdmin(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, videos, helper.DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) UpdateVideoStatus(ctx *gin.Context) {
	var req dto.UpdateVideoStatusRequestParams

	if err := ctx.ShouldBindQuery(&req); helper.CheckError(err, ctx) {
		return
	}

	err := adminCon.adminServ.UpdateVideoStatus(req)
	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("video status has been "+helper.UPDATE_SUCCESS, helper.EmptyObj{}, helper.DATA)
	ctx.JSON(http.StatusAccepted, response)
}

func (adminCon *adminController) FetchVerifyVideos(ctx *gin.Context) {
	var req dto.FetchVerifyVideosRequestParams

	if err := ctx.ShouldBindUri(&req); helper.CheckError(err, ctx) {
		return
	}

	videos, err := adminCon.adminServ.FetchVerifyVideos(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, videos, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}