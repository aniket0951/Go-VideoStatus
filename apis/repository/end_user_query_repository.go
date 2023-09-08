package repository

import (
	"errors"

	"github.com/aniket0951/video_status/apis/helper"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

func (eUserRepo *endUserRepository) CreateEndUser(args db.CreateAdminUserParams) (db.Users, error) {
	ctx, cancel := eUserRepo.Init()
	defer cancel()

	result, err := eUserRepo.db.Queries.CreateAdminUser(ctx, args)

	err = helper.HandleDBErr(err)
	if (result == db.Users{}) {
		return db.Users{}, errors.New("failed to create a account")
	}
	return result, err
}
