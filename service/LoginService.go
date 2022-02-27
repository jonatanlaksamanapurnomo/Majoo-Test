package services

import (
	"encoding/json"
	"majoo-be-test/enums"
	"majoo-be-test/functions/login"
	"majoo-be-test/response_models"
	"majoo-be-test/util"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload = response_models.GetLoginRequest()
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if payload.Username == "" || payload.Password == "" {
		http.Error(w, enums.INPUT_ERROR, enums.BAD_REQUEST)
		return
	}
	user, err := login.LoginFunction(payload, dbconn)
	if err == nil {
		token := util.GenerateJWT(w, user)
		var resp response_models.LoginResponse
		resp.Token = token
		resp.Message = enums.SUCCSES
		json.NewEncoder(w).Encode(&resp)
	} else {
		http.Error(w, err.Error(), enums.BAD_REQUEST)
	}
}
