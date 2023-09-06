package services

import "github.com/aniket0951/video_status/apis/repository"

type EndUserService interface{}

type endUserService struct {
	endUserRepo repository.EndUserRespository
}

func NewEndUserService(repo repository.EndUserRespository) EndUserService {
	return &endUserService{
		endUserRepo: repo,
	}
}
