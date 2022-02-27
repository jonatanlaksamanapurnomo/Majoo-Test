package login

import (
	"github.com/jinzhu/gorm"
	"majoo-be-test/models"
	"majoo-be-test/response_models"
)

func LoginFunction(payload response_models.LoginRequest, dbconn *gorm.DB) (models.User, error) {
	var user = models.GetUser()
	err := dbconn.Where("user_name = ? and password = ? ",
		payload.Username,
		models.GetMD5Hash(payload.Password)).Find(&user).Error
	return user, err
}
