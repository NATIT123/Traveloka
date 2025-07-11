package storage

import (
	"Traveloka/common"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
)

func (sql *sqlStore) UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *models.UserUpdate) error {

	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}
