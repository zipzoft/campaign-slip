package models

import "time"

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
		TurnOver      int           `json:"turn_over"bson:"turn_over"`
		TotalWithdraw int           `json:"total_withdraw"bson:"total_withdraw"`
		Wallet        int           `json:"wallet"bson:"wallet"`
		Point         int           `json:"point"bson:"point"`
		WalletSpin    int           `json:"wallet_spin"bson:"wallet_spin"`
		WalletPost    int           `json:"wallet_post"bson:"wallet_post"`
		RewardPoint   interface{}   `json:"reward_point"bson:"reward_point"`
		Under         []interface{} `json:"under"bson:"under"`
		DepositToday  struct {
			TopUp        int `json:"top_up"bson:"top_up"`
			Bonus        int `json:"bonus"bson:"bonus"`
			Count        int `json:"count"bson:"count"`
			Transactions []struct {
				BankName        string  `json:"bank_name" bson:"bank_name"`
				DateBank        string  `json:"date_bank"bson:"date_bank"`
				AccountTransfer string  `json:"account_transfer"bson:"account_transfer"`
				Amount          int     `json:"amount"bson:"amount"`
				Bonus           int     `json:"bonus"bson:"bonus"`
				TopUp           int     `json:"top_up"bson:"top_up"`
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

type Transaction struct {
	Username string `json:"username"bson:"username"`
	Detail   []struct {
		BankName        string  `json:"bank_name" bson:"bank_name"`
		DateBank        string  `json:"date_bank"bson:"date_bank"`
		AccountTransfer string  `json:"account_transfer"bson:"account_transfer"`
		Amount          int     `json:"amount"bson:"amount"`
		Bonus           int     `json:"bonus"bson:"bonus"`
		TopUp           int     `json:"top_up"bson:"top_up"`
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
}
