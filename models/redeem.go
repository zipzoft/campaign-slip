package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TransactionRedeem struct {
	ID           primitive.ObjectID `json:"id"bson:"_id,omitempty"`
	Username     string             `json:"username"bson:"username"`
	Prefix       string             `json:"prefix"bson:"prefix"`
	DateBank     string             `json:"date_bank"bson:"date_bank"`
	TopUp        int                `json:"top_up"bson:"top_up"`
	BeforeAmount float64            `json:"before_amount"bson:"before_amount"`
	AfterAmount  float64            `json:"after_amount"bson:"after_amount"`
	SlipNumber   int                `json:"slip_number"bson:"slip_number"`
	Coin         int                `json:"coin"bson:"coin"`
	CreatedAt    time.Time          `json:"created_at"bson:"created_at"`
	IsRedeem     bool               `json:"is_redeem"bson:"is_redeem"`
}
