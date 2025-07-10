package ginuser

import (
	"Traveloka/common"
	biz "Traveloka/modules/user/biz"
	storage "Traveloka/modules/user/storage/postgreSQL"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		// go func() {
		// 	defer common.Recovery()

		// 	var a []int
		// 	log.Println(a[0])
		// }()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// panic(common.ErrInvalidRequest(err))
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		bussiness := biz.NewGetUserBiz(store)

		data, err := bussiness.GetUserById(c.Request.Context(), id)

		if err != nil {
			// panic(err) // not best parctice
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
