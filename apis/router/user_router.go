package router

import (
	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/controllers"
	"github.com/aniket0951/video_status/apis/middleware"
	"github.com/aniket0951/video_status/apis/repository"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/gin-gonic/gin"
)

var JwtServ = services.NewJWTService()

func UserRoute(router *gin.Engine, store *apis.Store) {
	var (
		userRepo = repository.NewUserRepository(store)

		userServ = services.NewUserService(userRepo, JwtServ)
		userCont = controllers.NewUserController(userServ)
	)

	userauthroute := router.Group("/api")
	{
		userauthroute.POST("/create-admin-user", userCont.CreateAdminUser)
		userauthroute.GET("/login-admin", userCont.LoginAdminUser)
		userauthroute.PUT("/update-account-status", userCont.UpdateUserAccountStatus)
	}

	user_route := router.Group("/api", middleware.AuthorizeJWT(JwtServ))
	{
		user_route.GET("/admin-users/:page_id/:page_size", userCont.FetchAllUsers)
	}

}
