package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type WalletRequest struct {
	Prefix    string `json:"prefix" bson:"prefix" validate:"required"`
	Username  string `json:"username" bson:"username" validate:"required"`
	SettingID string `json:"setting_id" bson:"setting_id"`
	Note      string `json:"note" bson:"note"`
	Name      string `json:"name"bson:"name"`
}
type Wallet struct {
	Data []struct {
		Id                      string    `json:"id"`
		Company                 string    `json:"company"`
		Prefix                  string    `json:"prefix"`
		Name                    string    `json:"name"`
		Coin                    string    `json:"coin"`
		Period                  string    `json:"period"`
		Amount                  string    `json:"amount"`
		SpecialCoin             string    `json:"special_coin"`
		PeriodSpecialCoin       string    `json:"period_special_coin"`
		AmountPeriodSpecialCoin string    `json:"amount_period_special_coin"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
	} `json:"data"`
	Page  int `json:"page"`
	Total int `json:"total"`
}
type TransactionWallet struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StatusCode   int                `json:"status_code" bson:"status_code"`
	ReceivedDate time.Time          `json:"received_date" bson:"received_date"`
	Response     interface{}        `json:"response" bson:"response"`
}
