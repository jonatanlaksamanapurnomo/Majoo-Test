package models

type Transaction struct {
	ID           uint
	MerchantName string `gorm:"type:string" json:"merchant_name"`
	//Merchant   Merchant `gorm:"foreignKey:MerchantRefer"`
	//OutletId int `gorm:"type:int" json:"outlet_id"`
	//Outlet     Outlet   `gorm:"foreignKey:OutletRefer"`
	BillTotal float64 `gorm:"type:float" json:"bill_total"`
	//CreatedAt time.Time
	//CreatedBy int `gorm:"type:int" json:"created_by"`
	//UpdatedAt time.Time
	//UpdatedBy int `gorm:"type:int" json:"updated_by"`
}

func GetTransaction() Transaction {
	var transaction Transaction
	return transaction
}
func GetTransactions() []Transaction {
	var transaction []Transaction
	return transaction
}
