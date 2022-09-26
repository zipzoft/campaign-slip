package routers

import (
	"campiagn-slip/internal/controller"
	"campiagn-slip/internal/repository"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	transactionController := controller.NewTransactionController(repository.TransactionRepo{})
	settingController := controller.NewSettingController(repository.SettingRepo{})
	redeemController := controller.NewRedeemController(repository.RedeemRepo{})

	v1 := route.Group("api/v1")
	{
		v1.GET("transaction/topup", transactionController.GetTransaction) // ดึง API customer/top-up เพิ่มข้อมูล user_redeem  param query username,prefix
		v1.GET("settings", settingController.Condition)                   // get condition
		v1.POST("settings", settingController.InsertAndUpdateCondition)   // insert condition
		v1.PATCH("settings", settingController.InsertAndUpdateCondition)  // update condition
		v1.DELETE("settings", settingController.InsertAndUpdateCondition) // delete condition
		v1.POST("redeem", redeemController.Redeem)

		// ------------------------------------------------------------
		// Don't remove this line if you don't want to be maintainer.
		v1.GET("health-check", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}
}
