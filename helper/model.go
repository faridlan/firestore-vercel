package helper

import (
	"github.com/faridlan/firestore-vercel/model/domain"
	"github.com/faridlan/firestore-vercel/model/web"
)

func ToUserResponse(user *domain.User) *web.UserWeb {
	return &web.UserWeb{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
}

func ToUserResponses(users []*domain.User) []*web.UserWeb {

	userResponses := []*web.UserWeb{}

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses

}
