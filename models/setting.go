package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConditionTopUp struct {
	ID         primitive.ObjectID `json:"id"bson:"_id,omitempty"`
	Prefix     string             `json:"prefix"bson:"prefix"`
	MinTopUp   float64            `json:"min_top_up"bson:"min_top_up"`
	MaxBalance float64            `json:"max_balance"bson:"max_balance"`
	CreatedAt  time.Time          `json:"created_at"bson:"created_at"`
}

type ConditionRedeem struct {
	ID         primitive.ObjectID `json:"id"bson:"_id,omitempty"`
	Prefix     string             `json:"prefix"bson:"prefix"`
	SlipNumber int                `json:"slip_number"bson:"slip_number"`
	RedeemCoin int                `json:"redeem_coin"bson:"redeem_coin"`
	CreatedAt  time.Time          `json:"created_at"bson:"created_at"`
}

type Condition struct {
	ID     primitive.ObjectID `json:"id"bson:"_id,omitempty"`
	Prefix string             `json:"prefix"bson:"prefix"`
	Detail []struct {
		SlipNumber int `json:"slip_number"bson:"slip_number"`
		RedeemCoin int `json:"redeem_coin"bson:"redeem_coin"`
	} `json:"detail"bson:"detail"`
	MinTopUp   float64   `json:"min_top_up"bson:"min_top_up"`
	MaxBalance float64   `json:"max_balance"bson:"max_balance"`
	CreatedAt  time.Time `json:"created_at"bson:"created_at"`
}
