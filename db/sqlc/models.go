// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entry struct {
	Code      int64 `json:"code"`
	AccountID int64 `json:"accountID"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transfer struct {
	ID            int64         `json:"id"`
	FromAccountID sql.NullInt64 `json:"fromAccountID"`
	ToAccountID   sql.NullInt64 `json:"toAccountID"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}