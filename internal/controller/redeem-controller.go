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
	username := c.Query("username")
	prefix := c.Query("prefix")
	campaign := c.Query("campaign")
	trans := repository.TransactionRepo{}
	walletRequest, err, validateErr := trans.WalletValidate(username, prefix, campaign)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, validateErr)
		return
	}
	var userRedeem models.TransactionRedeem
	body, err := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &userRedeem)
	userRedeem, err = ctrl.repo.UpdateRedeem(userRedeem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	transaction, err := ctrl.repo.EarnCoin(walletRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = ctrl.repo.AddTransactionWallet(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"transaction": transaction, "user_redeem": userRedeem}, "message": "success"})
}
