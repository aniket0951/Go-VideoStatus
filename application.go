package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/video_status?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal(err)
	}
	store := apis.NewStore(conn)
	router_gin := gin.Default()

	router_gin.Use(cors.New(CORSConfig()))

	router_gin.Static("static", "static")

	router_gin.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "Application run success"})
	})

	router.UserRoute(router_gin, store)
	router.AdminRouter(router_gin, store)
	router_gin.Run(address)
}
