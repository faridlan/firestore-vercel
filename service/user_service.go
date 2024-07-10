package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-vercel/helper"
	"github.com/faridlan/firestore-vercel/model/domain"
	"github.com/faridlan/firestore-vercel/model/web"
	"github.com/faridlan/firestore-vercel/repository"
)

type UserService interface {
	Save(ctx context.Context, request *web.UserWeb) (*web.UserWeb, error)
	Find(ctx context.Context) ([]*web.UserWeb, error)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	Client   *firestore.Client
}

func NewUserService(userRepo repository.UserRepository, client *firestore.Client) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		Client:   client,
	}
}

func (service *UserServiceImpl) Save(ctx context.Context, request *web.UserWeb) (*web.UserWeb, error) {

	user := &domain.User{
		Name: request.Name,
		Age:  request.Age,
	}

	userResponse, err := service.UserRepo.Save(ctx, service.Client, user)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponse(userResponse), nil

}

func (service *UserServiceImpl) Find(ctx context.Context) ([]*web.UserWeb, error) {

	userResponse, err := service.UserRepo.Find(ctx, service.Client)
	if err != nil {
		return nil, err
	}

	return helper.ToUserResponses(userResponse), nil

}
