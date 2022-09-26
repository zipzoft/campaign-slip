package models

import (
	"time"
)

type TopUp struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    struct {
		Username      string        `json:"username"bson:"username"`
		FullName      string        `json:"full_name"bson:"full_name"`
		FirstName     string        `json:"first_name"bson:"first_name"`
		LastName      string        `json:"last_name"bson:"last_name"`
		Mobile        string        `json:"mobile"bson:"mobile"`
		AccountNumber string        `json:"account_number"bson:"account_number"`
		BankName      string        `json:"bank_name"bson:"bank_name"`
		Credit        float64       `json:"credit"bson:"credit"`
		TurnOver      float32       `json:"turn_over"bson:"turn_over"`
		TotalWithdraw float32       `json:"total_withdraw"bson:"total_withdraw"`
		Wallet        float32       `json:"wallet"bson:"wallet"`
		Point         float32       `json:"point"bson:"point"`
		WalletSpin    float32       `json:"wallet_spin"bson:"wallet_spin"`
		WalletPost    float32       `json:"wallet_post"bson:"wallet_post"`
		RewardPoint   interface{}   `json:"reward_point"bson:"reward_point"`
		Under         []interface{} `json:"under"bson:"under"`
		DepositToday  struct {
			TopUp        float64 `json:"top_up"bson:"top_up"`
			Bonus        float64 `json:"bonus"bson:"bonus"`
			Count        float64 `json:"count"bson:"count"`
			Transactions []struct {
				BankName        string  `json:"bank_name" bson:"bank_name"`
				DateBank        string  `json:"date_bank"bson:"date_bank"`
				AccountTransfer string  `json:"account_transfer"bson:"account_transfer"`
				Amount          float64 `json:"amount"bson:"amount"`
				Bonus           float64 `json:"bonus"bson:"bonus"`
				TopUp           float64 `json:"top_up"bson:"top_up"`
				BeforeAmount    float64 `json:"before_amount"bson:"before_amount"`
				AfterAmount     float64 `json:"after_amount"bson:"after_amount"`
				Remark          string  `json:"remark"bson:"remark"`
				ActionBy        struct {
					Name      interface{} `json:"name"bson:"name"`
					CreatedAt interface{} `json:"created_at"bson:"createdAt"`
					UpdatedAt interface{} `json:"updated_at"bson:"updatedAt"`
				} `json:"action_by"`
				CreatedAt time.Time `json:"created_at"bson:"created_at"`
				UpdatedAt time.Time `json:"updated_at"bson:"updated_at"`
			} `json:"transactions"bson:"transactions"`
		} `json:"deposit_today"bson:"deposit_today"`
	} `json:"data"bson:"data"`
}

type TransactionTopUp struct {
	Username string `json:"username"bson:"username"`
	Detail   []struct {
		BankName        string  `json:"bank_name" bson:"bank_name"`
		DateBank        string  `json:"date_bank"bson:"date_bank"`
		AccountTransfer string  `json:"account_transfer"bson:"account_transfer"`
		Amount          float64 `json:"amount"bson:"amount"`
		Bonus           float64 `json:"bonus"bson:"bonus"`
		TopUp           float64 `json:"top_up"bson:"top_up"`
		BeforeAmount    float64 `json:"before_amount"bson:"before_amount"`
		AfterAmount     float64 `json:"after_amount"bson:"after_amount"`
		Remark          string  `json:"remark"bson:"remark"`
		ActionBy        struct {
			Name      interface{} `json:"name"bson:"name"`
			CreatedAt interface{} `json:"created_at"bson:"createdAt"`
			UpdatedAt interface{} `json:"updated_at"bson:"updatedAt"`
		} `json:"action_by"`
		CreatedAt time.Time `json:"created_at"bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at"bson:"updated_at"`
	} `json:"detail"bson:"detail"`
	Page  int `json:"page"bson:"page"`
	Limit int `json:"limit"bson:"limit"`
	Total int `json:"total"bson:"total"`
}
