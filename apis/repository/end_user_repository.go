package repository

import (
	"context"
	"time"

	"github.com/aniket0951/video_status/apis"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

type EndUserRespository interface {
	Init() (context.Context, context.CancelFunc)
	CreateEndUser(args db.CreateAdminUserParams) (db.Users, error)
}

type endUserRepository struct {
	db *apis.Store
}

func NewEndUserRespository(db *apis.Store) EndUserRespository {
	return &endUserRepository{db: db}
}

func (eUserRepo *endUserRepository) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
