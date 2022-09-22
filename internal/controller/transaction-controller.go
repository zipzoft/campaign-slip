package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController struct {
	repo repository.TransactionRepository
}

func NewTransactionController(repo repository.TransactionRepository) *TransactionController {
	return &TransactionController{repo: repo}
}

func (ctrl *TransactionController) GetTransaction(c *gin.Context) {
	customer, err := ctrl.repo.GetCustomer(c)
	settingRepo := repository.SettingRepo{}
	redeemRepo := repository.RedeemRepo{}
	var transactionBonus models.TransactionTopUp
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transaction, err := ctrl.repo.GetTransaction(customer.Data.Username, customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//transaction := models.TransactionTopUp{}
	//body, err := io.ReadAll(c.Request.Body)
	//json.Unmarshal(body, &transaction)

	condition, err := settingRepo.FindOneCondition(customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	for _, m := range transaction.Detail {
		if m.BankName != "ADD_CREDIT" && m.BeforeAmount <= condition.MaxBalance && m.TopUp >= int(condition.MinTopUp) {
			transactionBonus.Username = customer.Data.Username
			transactionBonus.Detail = append(transactionBonus.Detail, m)
		}
	}
	if len(transactionBonus.Detail) >= condition.Detail[0].SlipNumber {
		result, err := ctrl.repo.InsertUserRedeem(transactionBonus, condition)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
		c.JSON(http.StatusOK,
			gin.H{"data": gin.H{"transaction": transaction, "user_redeem": userRedeem},
				"message": "เพิ่มข้อมูลสลิปที่ตรงตามเงื่อนไข" + strconv.Itoa(len(result)) + "สลิป"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": err.Error()}})
		return
	}
	userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"transaction": transaction, "user_redeem": userRedeem}, "message": "ไม่พบสลิปที่ตรงตามเงื่อนไขเพิ่มเติม"})
	return

}
