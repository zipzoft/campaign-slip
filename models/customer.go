package models

import "time"

type Customer struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Username        string      `json:"username"`
		Password        string      `json:"password"`
		Prefix          string      `json:"prefix"`
		FirstName       string      `json:"firstName"`
		LastName        string      `json:"lastName"`
		FirstNameEN     string      `json:"firstNameEN"`
		LastNameEN      string      `json:"lastNameEN"`
		AccountNumber   string      `json:"accountNumber"`
		BankName        string      `json:"bankName"`
		Tel             string      `json:"tel"`
		Line            string      `json:"line"`
		Under           interface{} `json:"under"`
		Recommend       string      `json:"recommend"`
		RecommendOther  string      `json:"recommendOther"`
		Bonus           string      `json:"bonus"`
		NotBonus        bool        `json:"notBonus"`
		AutoApprove     bool        `json:"autoApprove"`
		IsAuto          bool        `json:"isAuto"`
		CreateDate      *time.Time  `json:"create_date"`
		UpdateDate      *time.Time  `json:"update_date"`
		FirstTopupDate  *time.Time  `json:"first_topup_date"`
		LastTopupDate   *time.Time  `json:"last_topup_date"`
		LastDepositDate *time.Time  `json:"last_deposit_date"`
	} `json:"data"`
}
