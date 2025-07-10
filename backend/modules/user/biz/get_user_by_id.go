package biz

import (
	"Traveloka/common"
	"Traveloka/modules/user/models/postgreSQL"
	"context"
)

type GetUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*models.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*models.User, error) {

	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	return data, nil
}
