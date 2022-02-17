package restaurantlikemodel

type Like struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    int `json:"created_at" gorm:"column:created_at;"`
}

func (Like) TableName() string { return "restaurant_likes" }