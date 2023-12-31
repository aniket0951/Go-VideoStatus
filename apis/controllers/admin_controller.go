package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminController interface {
	UploadVideoByAdmin(*gin.Context)
	GetVideoByAdmin(*gin.Context)
	UpdateVideoStatus(*gin.Context)
	FetchVerifyVideos(*gin.Context)
	PublishedVideo(*gin.Context)
	FetchAllPublishedVideos(*gin.Context)
	UnPublishVideo(*gin.Context)
	FetchAllUnPublishVideo(*gin.Context)
	MakeVerificationFailed(*gin.Context)
	MakeUnPublishedVideo(*gin.Context)
	FetchAllVerificationFailedVideos(*gin.Context)
	FetchVerifyVideoFullDetails(*gin.Context)
	FetchVideoByAdminFullDetails(ctx *gin.Context)
	FetchPublishVideoFullDetails(ctx *gin.Context)
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
	status := ctx.PostForm("status")

	if title == "" || status == "" {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	admin_user, err := uuid.Parse(helper.TOKEN_ID)

	if err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusInternalServerError, response)
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

func (adminCon *adminController) PublishedVideo(ctx *gin.Context) {
	var req dto.PublishedVideoRequestParams

	if err := ctx.ShouldBindUri(&req); helper.CheckError(err, ctx) {
		return
	}

	err := adminCon.adminServ.PublishVideo(req)

	if err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.VIDEO_DATA)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildSuccessResponse("video has been published", helper.EmptyObj{}, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) FetchAllPublishedVideos(ctx *gin.Context) {
	var req dto.FetchVerifyVideosRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	videos, err := adminCon.adminServ.FetchAllPublishedVideos(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, videos, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) UnPublishVideo(ctx *gin.Context) {
	var req dto.PublishedVideoRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := adminCon.adminServ.UnPublishVideo(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("video has been unpublished", helper.EmptyObj{}, helper.VIDEO_DATA)
	ctx.JSON(http.StatusAccepted, response)
}

func (adminCon *adminController) FetchAllUnPublishVideo(ctx *gin.Context) {
	var req dto.FetchVerifyVideosRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	videos, err := adminCon.adminServ.FetchUnPublishVideos(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, videos, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) MakeVerificationFailed(ctx *gin.Context) {
	// verification_failed_by id taking as a current login user

	var req dto.CreateVerificationFailedRequestParam

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := adminCon.adminServ.MakeVerificationFailed(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("video verification failed has been successfully", helper.EmptyObj{}, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) MakeUnPublishedVideo(ctx *gin.Context) {
	// unpublish_by id taking as a current login user
	var req dto.CreateVerificationFailedRequestParam

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := adminCon.adminServ.MakeUnPublishedVideo(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse("video unpublish has been successfully", helper.EmptyObj{}, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) FetchAllVerificationFailedVideos(ctx *gin.Context) {
	var req dto.FetchVerifyVideosRequestParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		response := helper.BuildFailedResponse(helper.FETCHED_FAILED, err.Error(), helper.EmptyObj{}, helper.DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	videos, err := adminCon.adminServ.FetchAllVerificationFailedVideos(req)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, videos, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) FetchVerifyVideoFullDetails(ctx *gin.Context) {
	var video_id = ctx.Param("video_id")

	fmt.Println("video id : ", video_id)
	if video_id == "" {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	video_uuid, err := uuid.Parse(video_id)

	if helper.CheckError(err, ctx) {
		return
	}

	video_detail, err := adminCon.adminServ.FetchVerifyVideoFullDetails(video_uuid)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, video_detail, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) FetchVideoByAdminFullDetails(ctx *gin.Context) {
	video_id := ctx.Param("video_id")

	if video_id == "" {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	video_id_obj, err := uuid.Parse(video_id)

	if helper.CheckError(err, ctx) {
		return
	}

	video_detail, err := adminCon.adminServ.FetchVideoByAdminFullDetails(video_id_obj)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, video_detail, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}

func (adminCon *adminController) FetchPublishVideoFullDetails(ctx *gin.Context) {
	video_id := ctx.Param("video_id")

	if video_id == "" {
		helper.RequestBodyEmptyResponse(ctx)
		return
	}

	video_id_obj, err := uuid.Parse(video_id)

	if helper.CheckError(err, ctx) {
		return
	}

	video_detail, err := adminCon.adminServ.FetchPublishVideoFullDetails(video_id_obj)

	if helper.CheckError(err, ctx) {
		return
	}

	response := helper.BuildSuccessResponse(helper.FETCHED_SUCCESS, video_detail, helper.VIDEO_DATA)
	ctx.JSON(http.StatusOK, response)
}
