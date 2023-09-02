package services

import (
	"database/sql"
	"errors"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/repository"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/aniket0951/video_status/utils"
	"github.com/google/uuid"
)

type UserService interface {
	CreateAdminUser(dto.CreateAdminRequestParams) (dto.GetAdminUser, error)
	LoginAdminUser(dto.LoginUserRequestParams) (dto.GetAdminUser, error)
	UpdateUserAccountStatus(user_id uuid.UUID, account_status bool) error
	FetchAllUsers(dto.GetUsersRequestParams) ([]dto.GetAdminUser, error)
}

type userService struct {
	userRepo   repository.UserRepository
	jwtService JWTService
}

func NewUserService(repo repository.UserRepository, jwtSer JWTService) UserService {
	return &userService{
		userRepo:   repo,
		jwtService: jwtSer,
	}
}

func (userser *userService) CreateAdminUser(req dto.CreateAdminRequestParams) (dto.GetAdminUser, error) {
	hash_password := utils.HasAndSalt([]byte(req.Password))

	user_to_create := db.CreateAdminUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Contact:  req.Contact,
		Password: hash_password,
		UserType: req.UserType,
		IsAccountActive: sql.NullBool{
			Bool:  *req.IsAccountActive,
			Valid: true,
		},
	}

	new_admin, err := userser.userRepo.CreateAdminUser(user_to_create)

	err = helper.HandleDBErr(err)

	return dto.GetAdminUser{
		Name:     new_admin.Name,
		Email:    new_admin.Email,
		Contact:  new_admin.Contact,
		UserType: new_admin.UserType,
	}, err
}
func (userser *userService) LoginAdminUser(req dto.LoginUserRequestParams) (dto.GetAdminUser, error) {
	user, err := userser.userRepo.UserByEmail(req.Email)

	if !utils.ComparePassword(user.Password, []byte(req.Password)) {
		return dto.GetAdminUser{}, errors.New("password not matched")
	}

	if !user.IsAccountActive.Bool {
		return dto.GetAdminUser{}, errors.New("your account has been deactivate please contact with admin user")
	}

	token := userser.jwtService.GenerateToken(user.ID.String(), user.UserType)

	return dto.GetAdminUser{
		Name:        user.Name,
		Email:       user.Email,
		Contact:     user.Contact,
		UserType:    user.UserType,
		AccessToken: token,
	}, err
}
func (userser *userService) UpdateUserAccountStatus(user_id uuid.UUID, account_status bool) error {
	args := db.UpdateUserAccountActiveParams{
		ID: user_id,
		IsAccountActive: sql.NullBool{
			Bool:  account_status,
			Valid: true,
		},
	}

	_, err := userser.userRepo.UpdateUserAccountStatus(args)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found to change the account status")
		}
		return err
	}

	err = helper.HandleDBErr(err)
	return err
}

func (userser *userService) FetchAllUsers(req dto.GetUsersRequestParams) ([]dto.GetAdminUser, error) {
	args := db.GetUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	result, err := userser.userRepo.FetchAllUsers(args)

	if err != nil {
		return nil, err
	}

	err = helper.HandleDBErr(err)

	users := make([]dto.GetAdminUser, len(result))

	for i, user := range result {
		users[i] = dto.GetAdminUser{
			Name:     user.Name,
			Email:    user.Email,
			Contact:  user.Contact,
			UserType: user.UserType,
		}
	}

	return users, err
}
