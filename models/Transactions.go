package models

import (
	mssql "github.com/denisenkom/go-mssqldb"
	"time"
)

type Transaction struct {
	ID         uint
	MerchantId int     `gorm:"type:int" json:"merchant_id"`
	OutletId   int     `gorm:"type:int" json:"outlet_id"`
	BillTotal  float64 `gorm:"type:float" json:"bill_total"`
	CreatedAt  time.Time
	CreatedBy  int `gorm:"type:int" json:"created_by"`
	UpdatedAt  time.Time
	UpdatedBy  int `gorm:"type:int" json:"updated_by"`
}
type Transactions struct {
	MerchantName mssql.VarChar `gorm:"type:varchar(255)" json:"merchant_name"`
	Omzet        float64       `gorm:"type:float" json:"omzet"`
}

func GetTransaction() Transaction {
	var transaction Transaction
	return transaction
}
func GetTransactions() []Transaction {
	var transaction []Transaction
	return transaction
}
func GetReports() []Transactions {
	var reports []Transactions
	return reports
}
