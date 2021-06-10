package models

import "time"

type Rate struct {
	ID        string    `json:"id" bson:"id"`
	Exchange  float64   `json:"exchange" bson:"exchange"`
	TradeType string    `json:"trade_type" bson:"trade_type"`
	Price     float64   `json:"price"`
	Ts        time.Time `json:"ts" bson:"ts"`
}
