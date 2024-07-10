package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/faridlan/firestore-vercel/model/domain"
	"google.golang.org/api/iterator"
)

type UserRepository interface {
	Save(ctx context.Context, client *firestore.Client, user *domain.User) (*domain.User, error)
	Find(ctx context.Context, client *firestore.Client) ([]*domain.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, client *firestore.Client, user *domain.User) (*domain.User, error) {

	docRef, _, err := client.Collection("users").Add(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create new doc : " + err.Error())
	}

	user.ID = docRef.ID

	return user, nil
}

func (repository *UserRepositoryImpl) Find(ctx context.Context, client *firestore.Client) ([]*domain.User, error) {

	iter := client.Collection("users").Documents(ctx)

	defer iter.Stop()

	users := []*domain.User{}

	for {

		docSnapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, errors.New("failed to iterate doc : " + err.Error())
		}

		user := &domain.User{}
		err = docSnapshot.DataTo(&user)
		if err != nil {
			return nil, errors.New("failed to decode doc : " + err.Error())
		}

		user.ID = docSnapshot.Ref.ID

		users = append(users, user)
	}

	return users, nil
}
