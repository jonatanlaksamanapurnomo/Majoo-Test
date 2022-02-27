package models

import (
	"github.com/jinzhu/gorm"
)

type Merchant struct {
	gorm.Model
	UserId       int    `gorm:"type:int" json:"user_id"`
	User         User   `gorm:"foreignKey:UserRefer"`
	MerchantName string `gorm:"type:varchar(45);" json:"merchant_name"`
	UpdatedBy    int    `gorm:"type:int" json:"updated_by"`
	CreatedBy    int    `gorm:"type:int" json:"created_by"`
}

func GetMerchant() Merchant {
	var merchant Merchant
	return merchant
}
func GetMerchants() []Merchant {
	var merchant []Merchant
	return merchant
}
