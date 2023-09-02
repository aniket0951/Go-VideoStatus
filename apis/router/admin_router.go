package router

import (
	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/controllers"
	"github.com/aniket0951/video_status/apis/middleware"
	"github.com/aniket0951/video_status/apis/repository"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, store *apis.Store) {

	var (
		adminRepo = repository.NewAdminRepository(store)
		adminServ = services.NewAdminService(adminRepo)
		adminCont = controllers.NewAdminController(adminServ)
	)

	admin_router := router.Group("/api", middleware.AuthorizeJWT(JwtServ))
	{
		admin_router.POST("/upload-video", adminCont.UploadVideoByAdmin)
		admin_router.GET("/admin-videos/:page_id/:page_size", adminCont.GetVideoByAdmin)
		admin_router.PUT("/admin-videos/update-status", adminCont.UpdateVideoStatus)
		admin_router.GET("/admin-videos/verify-videos/:page_id/:page_size", adminCont.FetchVerifyVideos)
		admin_router.POST("/admin-videos/publish-video/:video_id", adminCont.PublishedVideo)
	}

	published_video := router.Group("/api", middleware.AuthorizeJWT(JwtServ))
	{
		published_video.GET("/published-videos/:page_id/:page_size", adminCont.FetchAllPublishedVideos)
		published_video.PUT("/published-videos/unpublish-videos/:video_id", adminCont.UnPublishVideo)
	}
}
