package biz

import (
	"Traveloka/common"
	"Traveloka/modules/user/models/postgreSQL"
	"context"
	"strings"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *models.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *models.UserCreate) error {
	email := strings.TrimSpace(data.Email)

	if email == "" {
		return common.NewCustomError(models.ErrEmailExisted, "field not found", "ErrFieldNotFound")
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
