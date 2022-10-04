package models

import (
	"time"
)

type ReportData struct {
	Username string              `json:"username"bson:"_id"`
	Data     []TransactionReport `json:"data"bson:"data"`
	Count    int                 `json:"count"bson:"count"`
}

type TransactionReport struct {
	Prefix     string    `json:"prefix"bson:"prefix"`
	DateBank   string    `json:"date_bank"bson:"date_bank"`
	SlipNumber int       `json:"slip_number"bson:"slip_number"`
	Coin       int64     `json:"coin"bson:"coin"`
	CreatedAt  time.Time `json:"created_at"bson:"created_at"`
	IsRedeem   bool      `json:"is_redeem"bson:"is_redeem"`
}
