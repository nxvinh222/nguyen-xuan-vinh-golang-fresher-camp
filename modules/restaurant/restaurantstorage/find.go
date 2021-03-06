package restaurantstorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	err := db.Where(conditions).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, common.RecordNotFound
	}
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
