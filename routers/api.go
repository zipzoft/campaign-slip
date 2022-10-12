package routers

import (
	"campiagn-slip/internal/controller"
	"campiagn-slip/internal/repository"
	"campiagn-slip/pkg/recaptcha"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	transactionController := controller.NewTransactionController(repository.TransactionRepo{})
	settingController := controller.NewSettingController(repository.SettingRepo{})
	redeemController := controller.NewRedeemController(repository.RedeemRepo{})
	reportController := controller.NewReportController(repository.ReportRepo{})

	v1 := route.Group("api/v1")
	{
		v1.GET("transaction/topup", transactionController.GetTransaction)     // ดึง API customer/top-up เพิ่มข้อมูล user_redeem  param query username,prefix
		v1.GET("settings", settingController.Condition)                       // get condition
		v1.POST("settings", settingController.InsertAndUpdateCondition)       // insert condition
		v1.PATCH("settings/:id", settingController.InsertAndUpdateCondition)  // update condition
		v1.DELETE("settings/:id", settingController.InsertAndUpdateCondition) // delete condition

		v1.POST("redeem", recaptcha.NewGinHandler(), redeemController.Redeem) // update user_redeem, earn_coin , transaction earn_coin
		v1.GET("report", reportController.ReportUserRedeem)                   // query param prefix,date_from format "DD/MM/YYYY 00:00",date_to  "DD/MM/YYYY 23:59"
		v1.POST("maintenance", transactionController.TestMaintenance)
		// ------------------------------------------------------------
		// Don't remove this line if you don't want to be maintainer.
		v1.GET("health-check", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}
}
