package report

import (
	"errors"
	"github.com/jinzhu/gorm"
	"majoo-be-test/enums"
	"majoo-be-test/models"
	"majoo-be-test/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetTransactionReport(r *http.Request, dbconn *gorm.DB) ([]models.Transactions, error) {
	if r.Header.Get("Authorization") == "" {
		return nil, errors.New(enums.AUTH_ERROR)
	}
	var loginToken string = strings.Split(r.Header.Get("Authorization"), " ")[1]
	var jwtClaims, jwtErr = util.GetJWTValue(loginToken)
	if jwtErr != nil {
		return nil, errors.New(enums.AUTH_ERROR)
	}
	var loginId uint = jwtClaims.Id
	var transaction = models.GetReports()
	err := dbconn.Select("transactions.bill_total as omzet , merchants.merchant_name").
		Joins("join merchants on merchants.id = transactions.merchant_id").
		Where("merchants.user_id = ? and date_part('month', transactions.created_at) = ?", loginId, time.November).
		Where("merchants.id = transactions.outlet_id").
		Scopes(Paginate(r)).Find(&transaction).Error
	return transaction, err
}
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
