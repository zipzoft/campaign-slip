package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	times "campiagn-slip/pkg/time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type TransactionController struct {
	repo repository.TransactionRepository
}

func NewTransactionController(repo repository.TransactionRepository) *TransactionController {
	return &TransactionController{repo: repo}
}

func TestMaintenance() bool {
	now := times.InBKK()
	validate1 := now.Truncate(24 * time.Hour).Add(22 * time.Hour)
	validate2 := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
	if now.After(validate1) &&
		now.Before(validate2) {
		return true
	}
	//else if now.Before(now.Truncate(24 * time.Hour).Add(1 * time.Hour)) {
	//	return true
	//}
	return false
}

func (ctrl *TransactionController) GetTransaction(c *gin.Context) {

	if TestMaintenance() {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"message": "ปรับปรุงระบบ"}})
		return
	}

	customer, err := ctrl.repo.GetCustomer(c)
	settingRepo := repository.SettingRepo{}
	redeemRepo := repository.RedeemRepo{}
	var transactionBonus models.TransactionTopUp
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"data": "", "user_redeem": "", "message": err.Error()}})
		return
	}
	transaction, err := ctrl.repo.GetTransaction(customer.Data.Username, customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"transaction": "", "user_redeem": "", "message": err.Error()}})
		return
	}
	condition, err := settingRepo.FindOneCondition(customer.Data.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"transaction": "", "user_redeem": "", "message": "ไม่พบ Setting ของ Prefix นี้ในระบบ"}})
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
			c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"transaction": "", "user_redeem": "", "message": err.Error()}})
			return
		}
		userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"transaction": "", "user_redeem": "", "message": err.Error()}})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{"data": bson.M{"transaction": transactionBonus, "user_redeem": userRedeem, "message": "เพิ่มข้อมูลสลิปที่ตรงตามเงื่อนไข" + strconv.Itoa(len(result)) + "สลิป"}})
		return
	}

	userRedeem, err := redeemRepo.GetUserRedeem(customer.Data.Username)
	c.JSON(http.StatusOK, gin.H{"data": bson.M{"transaction": transactionBonus, "user_redeem": userRedeem, "message": "ไม่พบสลิปที่ตรงตามเงื่อนไขเพิ่มเติม"}})
	return

}
