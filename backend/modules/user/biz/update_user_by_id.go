package biz

import (
	"Traveloka/common"
	"Traveloka/modules/user/models/enum"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
	"errors"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*models.User, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.UserUpdate) error
}

type updateItemBiz struct {
	store     UpdateItemStorage
	requester common.Requester
}

func NewUpdateItemBiz(store UpdateItemStorage, requester common.Requester) *updateItemBiz {
	return &updateItemBiz{store: store, requester: requester}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *models.UserUpdate) error {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(models.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	if data.Status != nil && *data.Status == enum.UserStatusDeleted {
		return common.ErrEntityDeleted(models.EntityName, models.ErrUserDeleted)
	}

	//isAdmin
	isOwner := biz.requester.GetUserId() == data.Id

	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission("user does not have permission to update this item", errors.New("No permission"))
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}
