package biz

import (
	"Traveloka/common"
	status "Traveloka/modules/user/models/enum"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
	"errors"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*models.User, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store     DeleteItemStorage
	requester common.Requester
}

func NewDeleteItemBiz(store DeleteItemStorage, requester common.Requester) *deleteItemBiz {
	return &deleteItemBiz{store: store, requester: requester}
}

func (biz *deleteItemBiz) DeletetemById(ctx context.Context, id int) error {

	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(models.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(models.EntityName, err)
	}

	if data.Status != nil && *data.Status == status.UserStatusDeleted {
		return common.NewCustomError(models.ErrUserDeleted, "user has been deleted", "ErrItemDeleted")
	}

	isOwner := biz.requester.GetUserId() == data.Id

	if !isOwner && !common.IsAdmin(biz.requester) {
		return common.ErrNoPermission("user does not have permission to delete this item", errors.New("No permission"))
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(models.EntityName, err)
	}

	return nil
}
