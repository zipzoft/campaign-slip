package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"sort"
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
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": "ไม่พบ User ในระบบ"}})
		return
	}
	transaction, err := ctrl.repo.GetTransaction(customer.Data.Username, customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": "มีข้อผิดพลาดในการดึง top-up transaction"}})
		return
	}
	condition, err := settingRepo.FindOneCondition(customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": "ไม่พบ Setting ในระบบ"}})
		return
	}
	// sort by bank_date
	sort.SliceStable(transaction.Detail, func(i, j int) bool {
		return transaction.Detail[i].DateBank < transaction.Detail[j].DateBank
	})
	// check condition
	for _, m := range transaction.Detail {
		if m.BankName != "ADD_CREDIT" && m.BeforeAmount <= condition.MaxBalance && m.TopUp >= condition.MinTopUp && m.Bonus == 0 {
			transactionBonus.Username = customer.Data.Username
			transactionBonus.Detail = append(transactionBonus.Detail, m)
		}
		if len(transactionBonus.Detail) == condition.Detail[len(condition.Detail)-1].SlipNumber {
			break
		}
	}
	// check lowest index case
	if len(transactionBonus.Detail) >= condition.Detail[0].SlipNumber {
		result, err := ctrl.repo.InsertUserRedeem(transactionBonus, condition)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": "ไม่สามารถเพิ่มข้อมูลได้"}})
			return
		}
		userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"data": bson.M{"transaction": transactionBonus, "user_redeem": userRedeem, "message": "เพิ่มข้อมูลสลิปที่ตรงตามเงื่อนไข" + strconv.Itoa(len(result)) + "สลิป"}})
		return
	}

	userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
	c.JSON(http.StatusOK, gin.H{"data": bson.M{"data": transactionBonus, "user_redeem": userRedeem, "message": "ไม่พบสลิปที่ตรงตามเงื่อนไขเพิ่มเติม"}})
	return

}
