package repository

import (
	"context"

	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

func (userrepo *userRepository) CreateAdminUser(user db.CreateAdminUserParams) (models.Users, error) {
	new_user, err := userrepo.db.Queries.CreateAdminUser(context.Background(), user)
	return models.Users(new_user), err
}

func (userrepo *userRepository) UserByEmail(user_mail string) (models.Users, error) {
	ctx, cancel := userrepo.Init()
	defer cancel()

	user, err := userrepo.db.Queries.GetUserByEmail(ctx, user_mail)
	return models.Users(user), err
}

func (userrepo *userRepository) UpdateUserAccountStatus(args db.UpdateUserAccountActiveParams) (models.Users, error) {
	ctx, cancel := userrepo.Init()
	defer cancel()

	user, err := userrepo.db.Queries.UpdateUserAccountActive(ctx, args)

	return models.Users(user), err
}

func (userrepo *userRepository) FetchAllUsers(args db.GetUsersParams) ([]models.Users, error) {
	ctx, cancel := userrepo.Init()
	defer cancel()

	result, err := userrepo.db.Queries.GetUsers(ctx, args)

	if err != nil {
		return nil, err
	}

	users := make([]models.Users, len(result))

	for i, user := range result {
		users[i] = models.Users(user)
	}

	return users, err
}
