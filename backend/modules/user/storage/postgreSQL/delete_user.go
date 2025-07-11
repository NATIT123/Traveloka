package storage

import (
	"Traveloka/common"
	"Traveloka/modules/user/models/enum"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
)

func (sql *sqlStore) DeleteUser(ctx context.Context, cond map[string]interface{}) error {

	if err := sql.db.Table(models.User{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": enum.UserStatusDeleted.String(),
		}).Error; err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
