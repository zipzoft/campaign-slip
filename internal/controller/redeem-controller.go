package controller

import (
	"campiagn-slip/internal/repository"
	"campiagn-slip/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"strconv"
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
		return
	}
	trans := repository.TransactionRepo{}

	walletRequest, err, validateErr := trans.WalletValidate(userRedeem.Username, campaign, userRedeem.Prefix, userRedeem.Coin)
	walletRequest.Coin = strconv.Itoa(int(userRedeem.Coin))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"message": "WallerValidate" + err.Error()}})
		return
	}
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, validateErr)
		return
	}
	_, err = ctrl.repo.UpdateRedeem(userRedeem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"message": "Update Redeem" + err.Error()}})
		return
	}
	transaction, err := ctrl.repo.EarnCoin(walletRequest)
	_, err = ctrl.repo.AddTransactionWallet(transaction, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": bson.M{"message": "Add TransactionWallet " + err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bson.M{"data": transaction.Response, "message": "success"}})
}
