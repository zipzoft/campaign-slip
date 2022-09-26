package repository

import (
	"campiagn-slip/config"
	"campiagn-slip/models"
	"campiagn-slip/pkg/database"
	times "campiagn-slip/pkg/time"
	"campiagn-slip/pkg/validator"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"strings"
	"time"
)

var _ TransactionRepository = (*TransactionRepo)(nil)

type TransactionRepository interface {
	GetTransaction(username, prefix string) (*models.TransactionTopUp, error)
	GetCustomer(c *gin.Context) (*models.Customer, error)
	InsertUserRedeem(transaction models.TransactionTopUp, condition models.Condition) ([]models.TransactionRedeem, error)
	WalletValidate(username, campaign, prefix string) (*models.WalletRequest, error, []validator.ApiError)
	GetSettingID(name, prefix string) (*models.Wallet, error)
}

type TransactionRepo struct {
	//
}

func (t TransactionRepo) GetTransaction(username, prefix string) (*models.TransactionTopUp, error) {
	conn := config.GetConfig()
	Username := strings.ToUpper(username)
	Prefix := strings.ToLower(prefix)

	request, err := http.NewRequest("GET", conn.AMMBOUrl+"/v1/customers/top-up", nil)

	if err != nil {
		return nil, err
	}

	//request.Header.Set("Accept", "application/json")
	query := request.URL.Query()
	query.Add("prefix", Prefix)
	query.Add("username", Username)
	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(request)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err != nil {
		return nil, err
	}

	var topUp models.TopUp

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &topUp); err != nil {
		return nil, err
	}

	transaction := models.TransactionTopUp{}
	for _, item := range topUp.Data.DepositToday.Transactions {
		transaction.Username = Username
		transaction.Detail = append(transaction.Detail, item)
	}

	return &transaction, nil

}
func (t TransactionRepo) GetCustomer(c *gin.Context) (*models.Customer, error) {

	conn := config.GetConfig()
	Username := strings.ToUpper(c.Query("username"))
	Prefix := c.Query("prefix")

	request, err := http.NewRequest("GET", conn.AMMBOUrl+"/v1/customers/detail", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	query := request.URL.Query()
	query.Add("prefix", Prefix)
	query.Add("username", Username)
	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(request)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err != nil {
		return nil, err
	}

	var customer *models.Customer

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (t TransactionRepo) InsertUserRedeem(transaction models.TransactionTopUp, condition models.Condition) ([]models.TransactionRedeem, error) {

	result := make([]models.TransactionRedeem, 0)
	model := models.TransactionRedeem{}

	con := 0

	for i, s := range transaction.Detail {
		if condition.Detail[con].SlipNumber == i+1 {
			model.ID = primitive.NewObjectID()
			model.Username = transaction.Username
			model.Prefix = strings.ToLower(condition.Prefix)
			model.DateBank = s.DateBank
			model.TopUp = s.TopUp
			model.BeforeAmount = s.BeforeAmount
			model.AfterAmount = s.AfterAmount
			model.SlipNumber = condition.Detail[con].SlipNumber
			model.Coin = condition.Detail[con].RedeemCoin
			model.CreatedAt = time.Now().Add(7 * time.Hour)
			model.IsRedeem = false
			con++
			filter := bson.M{
				"username":    model.Username,
				"date_bank":   model.DateBank,
				"slip_number": model.SlipNumber,
				"created_at": bson.M{
					"$gte": primitive.NewDateTimeFromTime(times.InBKK().Truncate(24 * time.Hour)),
				},
			}
			checkRedeem := models.TransactionRedeem{}
			err := database.FindOne("user_redeem", filter).Decode(&checkRedeem)
			if err != nil && err.Error() != "mongo: no documents in result" {
				return nil, err
			}
			if checkRedeem.Username == "" {
				_, err := database.InsertOne("user_redeem", model)
				if err != nil {
					return nil, err
				}
				result = append(result, model)
			}
		}

	}
	return result, nil
}
func (t TransactionRepo) WalletValidate(username, campaign, prefix string) (*models.WalletRequest, error, []validator.ApiError) {
	walletRequest := &models.WalletRequest{}
	wallet, err := t.GetSettingID(campaign, prefix)
	if err != nil {
		return nil, err, nil
	}
	if len(wallet.Data) > 0 {
		walletRequest.Username = username
		walletRequest.Prefix = prefix
		walletRequest.Name = campaign
		walletRequest.SettingID = wallet.Data[0].Id
		walletRequest.Note = walletRequest.Name
	} else {
		return nil, errors.New("ไม่พบข้อมูลกิจกรรม" + campaign), nil
	}

	validateErr := validator.Validate(walletRequest)

	if validateErr != nil {
		return nil, nil, validateErr
	}

	return walletRequest, nil, nil
}
func (t TransactionRepo) GetSettingID(name, prefix string) (*models.Wallet, error) {
	request, _ := http.NewRequest("GET", config.GetConfig().WalletAPI+"/api/v1/settings", nil)

	query := request.URL.Query()
	query.Add("name", name)
	query.Add("prefix", prefix)

	request.URL.RawQuery = query.Encode()
	request.Header.Add("Accept", "application/json")

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var result *models.Wallet

	err = json.NewDecoder(response.Body).Decode(&result)
	fmt.Println(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if response.StatusCode == 400 {
		return nil, errors.New("เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง")
	}

	return result, nil
}
