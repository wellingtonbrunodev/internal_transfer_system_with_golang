package model

type Account struct {
	ID      int64  `json:"account_id"`
	Balance string `json:"balance"`
}
