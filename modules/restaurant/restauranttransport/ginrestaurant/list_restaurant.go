package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurant/restaurantbiz"
	"food-delivery/modules/restaurant/restaurantmodel"
	"food-delivery/modules/restaurant/restaurantstorage"
	restaurantlikestorage "food-delivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		err := c.ShouldBind(&filter)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		err = c.ShouldBind(&paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
