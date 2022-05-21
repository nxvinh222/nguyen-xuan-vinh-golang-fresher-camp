package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurant/restaurantbiz"
	"food-delivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param(("id")))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		err = biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
