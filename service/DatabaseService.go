package services

import (
	"errors"
	"github.com/jinzhu/gorm"
	"majoo-be-test/models"
)

var dbconn *gorm.DB

func SetDB(db *gorm.DB) {
	dbconn = db
	var user = models.GetUser()
	var merchant = models.GetMerchant()
	var outlet = models.GetOutlet()
	var transaction = models.GetTransaction()
	dbconn.AutoMigrate(&user)
	dbconn.AutoMigrate(&merchant)
	dbconn.AutoMigrate(&outlet)
	dbconn.AutoMigrate(&transaction)
	//seeder if empty
	if err := dbconn.Find(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		firstUser := models.User{Name: "Admin 1", UserName: "admin1", Password: models.GetMD5Hash("admin1"), CreatedBy: 1, UpdatedBy: 1}
		secondUser := models.User{Name: "Admin 2", UserName: "admin2", Password: models.GetMD5Hash("admin2"), CreatedBy: 1, UpdatedBy: 1}
		dbconn.Create(&firstUser)
		dbconn.Create(&secondUser)
	}
	if err := dbconn.Find(&merchant).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		firstMerchant := models.Merchant{UserId: 1, MerchantName: "merchant 1", CreatedBy: 1, UpdatedBy: 1}
		secondMerchant := models.Merchant{UserId: 2, MerchantName: "merchant 2", CreatedBy: 2, UpdatedBy: 2}
		dbconn.Create(&firstMerchant)
		dbconn.Create(&secondMerchant)
	}
	if err := dbconn.Find(&outlet).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		firstOutlet := models.Outlet{MerchantId: 1, OutletName: "outlet 1", CreatedBy: 1, UpdatedBy: 1}
		secondOutlet := models.Outlet{MerchantId: 2, OutletName: "Outlet 1", CreatedBy: 2, UpdatedBy: 2}
		thridOutlet := models.Outlet{MerchantId: 1, OutletName: "Outlet 2", CreatedBy: 1, UpdatedBy: 1}
		dbconn.Create(&firstOutlet)
		dbconn.Create(&secondOutlet)
		dbconn.Create(&thridOutlet)
	}
	if err := dbconn.Find(&transaction).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		dbconn.Exec(" insert into Transactions values \n  (1, 1, 1, 2000, '2021-11-01 12:30:04', 1, '2021-11-01 12:30:04',1), \n  (2, 1, 1, 2500, '2021-11-01 17:20:14', 1, '2021-11-01 17:20:14',1),\n  (3, 1, 1, 4000, '2021-11-02 12:30:04', 1, '2021-11-02 12:30:04',1),\n  (4, 1, 1, 1000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1),\n  (5, 1, 1, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1),\n  (6, 1, 3, 2000, '2021-11-02 18:30:04', 1, '2021-11-02 18:30:04',1), \n  (7, 1, 3, 2500, '2021-11-03 17:20:14', 1, '2021-11-03 17:20:14',1),\n  (8, 1, 3, 4000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1),\n  (9, 1, 3, 1000, '2021-11-04 12:31:04', 1, '2021-11-04 12:31:04',1),\n  (10, 1, 3, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1),\n  (11, 2, 2, 2000, '2021-11-01 18:30:04', 2, '2021-11-01 18:30:04',2), \n  (12, 2, 2, 2500, '2021-11-02 17:20:14', 2, '2021-11-02 17:20:14',2),\n  (13, 2, 2, 4000, '2021-11-03 12:30:04', 2, '2021-11-03 12:30:04',2),\n  (14, 2, 2, 1000, '2021-11-04 12:31:04', 2, '2021-11-04 12:31:04',2),\n  (15, 2, 2, 7000, '2021-11-05 16:59:30', 2, '2021-11-05 16:59:30',2),\n  (16, 2, 2, 2000, '2021-11-05 18:30:04', 2, '2021-11-05 18:30:04',2), \n  (17, 2, 2, 2500, '2021-11-06 17:20:14', 2, '2021-11-06 17:20:14',2),\n  (18, 2, 2, 4000, '2021-11-07 12:30:04', 2, '2021-11-07 12:30:04',2),\n  (19, 2, 2, 1000, '2021-11-08 12:31:04', 2, '2021-11-08 12:31:04',2),\n  (20, 2, 2, 7000, '2021-11-09 16:59:30', 2, '2021-11-09 16:59:30',2),\n  (21, 2, 2, 1000, '2021-11-10 12:31:04', 2, '2021-11-10 12:31:04',2),\n  (22, 2, 2, 7000, '2021-11-11 16:59:30', 2, '2021-11-11 16:59:30',2);")
	}

}
