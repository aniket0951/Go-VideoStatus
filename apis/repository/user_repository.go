package repository

import (
	"context"
	"time"

	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

type UserRepository interface {
	CreateAdminUser(db.CreateAdminUserParams) (models.Users, error)
	UserByEmail(user_mail string) (models.Users, error)
	UpdateUserAccountStatus(args db.UpdateUserAccountActiveParams) (models.Users, error)
	FetchAllUsers(args db.GetUsersParams) ([]models.Users, error)
}

type userRepository struct {
	db *apis.Store
}

func NewUserRepository(db *apis.Store) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (userrepo *userRepository) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
