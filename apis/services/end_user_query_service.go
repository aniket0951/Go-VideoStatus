package services

import (
	"database/sql"

	"github.com/aniket0951/video_status/apis/dto"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

func (eUserSer *endUserService) CreateEndUser(req dto.CreateEndUserRequestParamDTO) (dto.GetAdminUser, error) {
	end_user_default := dto.CreateEndUser{}
	end_user_default.InitDefaultEndUser()
	end_user_default.Contact = req.Contact

	args := db.CreateAdminUserParams{
		Name:     end_user_default.Name,
		Email:    end_user_default.Email,
		Contact:  end_user_default.Contact,
		Password: "",
		UserType: end_user_default.UserType,
		IsAccountActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	user, err := eUserSer.endUserRepo.CreateEndUser(args)

	return dto.GetAdminUser{
		Name:     user.Name,
		Email:    user.Email,
		Contact:  user.Contact,
		UserType: user.UserType,
	}, err
}
