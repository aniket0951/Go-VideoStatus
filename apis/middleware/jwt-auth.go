package middleware

import (
	"fmt"
	"net/http"

	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.BuildFailedResponse(helper.FAILED_PROCESS, "Token not found", helper.EmptyObj{}, helper.USER_DATA)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if err != nil {
			response := helper.BuildFailedResponse("Invalid token provided !", err.Error(), helper.EmptyObj{}, helper.USER_DATA)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userId := fmt.Sprintf("%v", claims["user_id"])
			userType := fmt.Sprintf("%v", claims["user_type"])

			helper.TOKEN_ID = userId
			helper.USER_TYPE = userType

		}

	}
}
