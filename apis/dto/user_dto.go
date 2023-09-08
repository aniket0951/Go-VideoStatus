package dto

type CreateAdminRequestParams struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Contact         string `json:"contact" binding:"required,min=10"`
	Password        string `json:"password" binding:"required"`
	UserType        string `json:"user_type" binding:"required,oneof=ADMIN ENDUSER"`
	IsAccountActive *bool  `json:"is_account_active" binding:"required"`
}

type LoginUserRequestParams struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type UpdateUserAccountStatusRequestParams struct {
	Id            string `form:"id" binding:"required"`
	AccountStatus *bool  `form:"account_status" binding:"required"`
}

type GetUsersRequestParams struct {
	PageID   int32 `uri:"page_id"`
	PageSize int32 `uri:"page_size"`
}

type GetAdminUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Contact     string `json:"contact"`
	UserType    string `json:"user_type"`
	AccessToken string `json:"access_token,omitempty"`
}

// ---------------------------  END USER STRUCTS ----------------------- //
type CreateEndUserRequestParamDTO struct {
	Contact string `form:"contact" binding:"required"`
}
type CreateEndUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Contact     string `json:"contact"`
	UserType    string `json:"user_type"`
	AccessToken string `json:"access_token,omitempty"`
}

func (endUser *CreateEndUser) InitDefaultEndUser() {
	endUser.Name = "video_status_" + endUser.Contact
	endUser.Email = "video_status_@gmail.com"
	endUser.UserType = "END_USER"
}
