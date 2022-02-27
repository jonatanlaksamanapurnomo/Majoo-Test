package models

import (
	"github.com/jinzhu/gorm"
)

type Outlet struct {
	gorm.Model
	MerchantId int      `gorm:"type:int" json:"merchant_id"`
	Merchant   Merchant `gorm:"foreignKey:MerchantRefer"`
	OutletName string   `gorm:"type:varchar(45);" json:"outlet_name"`
	UpdatedBy  int      `gorm:"type:int" json:"updated_by"`
	CreatedBy  int      `gorm:"type:int" json:"created_by"`
}

func GetOutlet() Outlet {
	var outlet Outlet
	return outlet
}
