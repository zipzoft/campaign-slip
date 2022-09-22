package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type RedeemController struct {
	repo repository.RedeemRepository
}

func NewRedeemController(repo repository.RedeemRepository) *RedeemController {
	return &RedeemController{repo: repo}
}

func (ctrl *RedeemController) Redeem(c *gin.Context) {
	campaign := c.Query("campaign")
	var userRedeem models.TransactionRedeem
	body, err := io.ReadAll(c.Request.Body)
	err = json.Unmarshal(body, &userRedeem)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	trans := repository.TransactionRepo{}
	walletRequest, err, validateErr := trans.WalletValidate(userRedeem.Username, campaign, userRedeem.Prefix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, validateErr)
		return
	}
	_, err = ctrl.repo.UpdateRedeem(userRedeem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	transaction, err := ctrl.repo.EarnCoin(walletRequest)
	_, err = ctrl.repo.AddTransactionWallet(transaction, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "earn coin/transaction wallet invalid!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"transaction": transaction}, "message": "success"})
}
