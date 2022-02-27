package services

import (
	"encoding/json"
	"majoo-be-test/enums"
	"majoo-be-test/functions/report"
	"majoo-be-test/response_models"
	"net/http"
)

func Report(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transaction, err := report.GetTransactionReport(r, dbconn, w)
	if err == nil {
		var resp response_models.ReportResponse
		resp.Data = transaction
		resp.Message = enums.SUCCSES
		json.NewEncoder(w).Encode(&transaction)
	} else {
		http.Error(w, err.Error(), 400)
	}
}
