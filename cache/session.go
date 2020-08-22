package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/RajibDas-123/ms-grpc-auth/auth/model"

	"github.com/RajibDas-123/ms-grpc-auth/auth/logging"
)

//To check whether the session token is present in the cache or not
func CheckSessionToken(sessionToken string) bool {
	logging.CacheLogger.Infof("func:CheckSessionToken:Checking whether Session token is in session or not ", sessionToken)
	err := PubSub.Get(sessionToken).Err()
	if err != nil {
		logging.CacheLogger.Errorf("func:CheckSessionToken: Session Token  ", sessionToken, " not found.")
		return false
	} else {
		logging.CacheLogger.Infof("func:CheckSessionToken: Session Token  ", sessionToken, " found.")
		return true

	}
}

//To Fetch the data against the session token found in above function
func GetSessionToken(sessionToken string) (model.SessionData, bool) {
	var sessionData model.SessionData
	res, err := PubSub.Get(sessionToken).Result()
	if err != nil {
		logging.CacheLogger.Errorf("func:GetSessionToken: No data found against session token ", sessionToken)
		return sessionData, false
	} else {
		logging.CacheLogger.Infof("func:GetSessionToken: Data found against session token ", sessionToken)
		err := json.Unmarshal([]byte(res), &sessionData)
		if err != nil {
			logging.CacheLogger.Errorf("func:GetSessionToken: Unable to convert JSON to object with error ", err)
			return sessionData, false
		} else {
			logging.CacheLogger.Infof("func:GetSessionToken: Converted JSON to object.")
			return sessionData, true
		}
	}
}

//To set the session token and convert the session data to json format.
func SetSessionToken(role, email, sessionToken string, uid int32) bool {
	sessionData := model.SessionData{
		Role:  role,
		Email: email,
		UID:   strconv.Itoa(int(uid)),
	}
	res, err := json.Marshal(sessionData)
	if err != nil {
		logging.CacheLogger.Errorf("func:SetSessionToken: Unable to convert Session data to json format.")
		return false
	} else {
		err := PubSub.Set(sessionToken, res, 24*time.Hour).Err()
		if err != nil {
			logging.CacheLogger.Errorf("func:SetSessionToken: Unable to set Session token with error, ", err)
			return false
		} else {
			logging.CacheLogger.Infof("func:SetSessionToken: Session token ", sessionToken, " set successfuly")
			return true
		}
	}
}
