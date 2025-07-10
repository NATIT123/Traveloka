package storage

import (
	"Traveloka/common"
	models "Traveloka/modules/user/models/postgreSQL"
	"context"
)

func (sql *sqlStore) CreateUser(ctx context.Context, data *models.UserCreate) error {
	db := sql.db.Begin()
	if err := sql.db.Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
