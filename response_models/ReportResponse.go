package response_models

import "majoo-be-test/models"

type ReportResponse struct {
	Data    []models.Transaction `json:"data"`
	Message string               `json:"message"`
}
