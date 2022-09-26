package repository

import (
	"bytes"
	"campiagn-slip/config"
	"campiagn-slip/models"
	"campiagn-slip/pkg/database"
	times "campiagn-slip/pkg/time"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"time"
)

var _ RedeemRepository = (*RedeemRepo)(nil)

type RedeemRepository interface {
	UpdateRedeem(userRedeem models.TransactionRedeem) (*models.TransactionRedeem, error)
	EarnCoin(wallet *models.WalletRequest) (*models.TransactionWallet, error)
	AddTransactionWallet(transaction *models.TransactionWallet, err error) (*models.TransactionWallet, error)
	GetUserRedeem(username string) ([]models.TransactionRedeem, error)
}
type RedeemRepo struct {
	//
}

func (r RedeemRepo) UpdateRedeem(userRedeem models.TransactionRedeem) (*models.TransactionRedeem, error) {

	filter := bson.M{"_id": userRedeem.ID, "is_redeem": false}
	userRedeem.IsRedeem = true
	result, err := database.UpdateOne("user_redeem", filter, userRedeem)
	if result.ModifiedCount == 0 {
		return nil, errors.New("ModifiedCount = 0")
	}
	return &userRedeem, err

}
func (r RedeemRepo) EarnCoin(wallet *models.WalletRequest) (*models.TransactionWallet, error) {

	jsonValue, err := json.Marshal(wallet)

	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(jsonValue)

	req, err := http.NewRequest("POST", config.New().WalletAPI+"/api/v1/transactions", jsonBody)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	result := &models.TransactionWallet{}

	mapStruct := map[string]interface{}{}

	if resp.StatusCode == 400 {
		err = errors.New("เกิดข้อผิดพลาดกรุณาลองใหม่อีกครั้ง")
	}

	err = json.Unmarshal(response, &mapStruct)

	result.Response = mapStruct["data"]

	if err != nil {
		return nil, err
	}

	result.StatusCode = resp.StatusCode
	result.ReceivedDate = time.Now().Truncate(24 * time.Hour)

	return result, nil
}
func (r RedeemRepo) AddTransactionWallet(transaction *models.TransactionWallet, err error) (*models.TransactionWallet, error) {
	transaction.ID = primitive.NewObjectID()
	if err != nil {
		transaction.Response = map[string]interface{}{
			"message": "เกิดข้อผิดพลาดในการเพิ่ม coin",
		}
		_, err = database.InsertOne("transactions_wallet", transaction)
	}
	_, err = database.InsertOne("transactions_wallet", transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
func (r RedeemRepo) GetUserRedeem(username string) ([]models.TransactionRedeem, error) {

	userRedeem := make([]models.TransactionRedeem, 0)
	filter := bson.M{
		"username": username,
		"created_at": bson.M{
			"$gte": primitive.NewDateTimeFromTime(times.InBKK().Truncate(24 * time.Hour)),
		},
	}
	_, err := database.Find("user_redeem", filter, &userRedeem)
	if err != nil {
		return nil, err
	}
	return userRedeem, err
}
