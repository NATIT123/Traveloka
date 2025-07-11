package biz

import (
	"Traveloka/common"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
)

type GetUserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}, moreInfo ...string) (*models.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*models.User, error) {

	data, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	return data, nil
}
