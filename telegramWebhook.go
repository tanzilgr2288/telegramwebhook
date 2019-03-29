package telegramwebhook

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// @author Kevin Setiawan Tanzil [Tom Cruise]
// @desc this file for handler simple webhook bind & unbind to telegram API
// @file telegramWebhook.go
// @since 28 Maret 2019

// ReturnSetWebHookAndDelete for store and handle the return setWebHookInfo and Delete
type ReturnSetWebHookAndDelete struct {
	OK          bool   `json:"ok" bson:"ok" query:"ok" form:"ok"`
	Result      bool   `json:"result" bson:"result" query:"result" form:"result"`
	Description string `json:"description" bson:"description" query:"description" form:"description"`
	HTTPCODE    int
	HTTPMessage string
}

// ReturnGetWebHookInfo for store and handle the return getWebHookInfo
type ReturnGetWebHookInfo struct {
	OK          bool                 `json:"ok" bson:"ok" query:"ok" form:"ok"`
	Result      ResultGetWebhookInfo `json:"result" bson:"result" query:"result" form:"result"`
	HTTPCODE    int
	HTTPMessage string
}

// ResultGetWebhookInfo this struct just only for getwebhookinfo result, avoiding not using interface, so will be easily to use
type ResultGetWebhookInfo struct {
	URL                  string   `json:"url" bson:"url" query:"url" form:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate" bson:"has_custom_certificate" query:"has_custom_certificate" form:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count" bson:"pending_update_count" query:"pending_update_count" form:"pending_update_count"`
	MaxConnections       int      `json:"max_connections" bson:"max_connections" query:"max_connections" form:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates" bson:"allowed_updates" query:"allowed_updates" form:"allowed_updates"`
	LastErrorDate        int      `json:"last_error_date" bson:"last_error_date" query:"last_error_date" form:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message" bson:"last_error_message" query:"last_error_message" form:"last_error_message"`
}

// SetWebhook this function for setting webbook into telegram official API
func SetWebhook(callbackURL string, maxConnection int, allowUpdate []string, token string) (ReturnSetWebHookAndDelete, error) {

	// checking the maxConnection only 1 - 100 by default 40
	if maxConnection < 1 && maxConnection > 100 {
		maxConnection = 40
	}

	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/setWebhook"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)

	q := request.URL.Query()
	q.Add("url", callbackURL)
	q.Add("max_connections", strconv.Itoa(maxConnection))
	if allowUpdate != nil {
		for _, text := range allowUpdate {
			q.Add("allowed_updates", text)
		}
	}
	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSetWebHookAndDelete{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSetWebHookAndDelete
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// GetWebhookInfo this func for get webhook info from your telegram bot
func GetWebhookInfo(token string) (ReturnGetWebHookInfo, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/getWebhookInfo"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnGetWebHookInfo{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnGetWebHookInfo
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// DeleteWebHookInstance this func for delete webhook instance from telegram
func DeleteWebHookInstance(token string) (ReturnSetWebHookAndDelete, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/deleteWebhook"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSetWebHookAndDelete{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSetWebHookAndDelete
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// Comming Soon
// SetWebhookWithCertificate this function for setting webbook into telegram official API with Certificate
