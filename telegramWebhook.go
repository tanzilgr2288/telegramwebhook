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

const (
	// Constanta for parse style HTMLs
	PARSE_TYPE_HTML = "HTML"
	// Constanta for parse style markdown V2
	PARSE_TYPE_MARKDOWNV2 = "MarkdownV2"
)

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

// FromTelegram this sturct for handle from return from telegram
type FromTelegram struct {
	ID           int    `json:"id" bson:"id" query:"id" form:"id"`
	IsBot        bool   `json:"is_bot" bson:"is_bot" query:"is_bot" form:"is_bot"`
	FirstName    string `json:"first_name" bson:"first_name" query:"first_name" form:"first_name"`
	LastName     string `json:"last_name" bson:"last_name" query:"last_name" form:"last_name"`
	Username     string `json:"username" bson:"username" query:"username" form:"username"`
	LanguageCode string `json:"language_code" bson:"language_code" query:"language_code" form:"language_code"`
}

// ChatTelegram this struct for handler chat from telegram
type ChatTelegram struct {
	ID        int    `json:"id" bson:"id" query:"id" form:"id"`
	FirstName string `json:"first_name" bson:"first_name" query:"first_name" form:"first_name"`
	LastName  string `json:"last_name" bson:"last_name" query:"last_name" form:"last_name"`
	Username  string `json:"username" bson:"username" query:"username" form:"username"`
	Type      string `json:"type" bson:"type" query:"type" form:"type"`
}

// ReplyToMessage this struct for handler ReplyMessage from telegram
type ReplyToMessage struct {
	MessageID int          `json:"message_id" bson:"message_id" query:"message_id" form:"message_id"`
	From      FromTelegram `json:"from" bson:"from" query:"from" form:"from"`
	Chat      ChatTelegram `json:"chat" bson:"chat" query:"chat" form:"chat"`
	Date      int          `json:"date" bson:"date" query:"date" form:"date"`
	Text      string       `json:"text" bson:"text" query:"text" form:"text"`
}

// ResultSendMessage this struct for handler the result send message with reply Mode and normal
type ResultSendMessage struct {
	MessageID      int            `json:"message_id" bson:"message_id" query:"message_id" form:"message_id"`
	From           FromTelegram   `json:"from" bson:"from" query:"from" form:"from"`
	Chat           ChatTelegram   `json:"chat" bson:"chat" query:"chat" form:"chat"`
	Date           string         `json:"date" bson:"date" query:"date" form:"date"`
	Text           string         `json:"text" bson:"text" query:"text" form:"text"`
	ReplyToMessage ReplyToMessage `json:"reply_to_message" bson:"reply_to_message" query:"reply_to_message" form:"reply_to_message"`
}

// ReturnSendMessage this struct for handler the return send message with reply Mode and normal from Telegram API
type ReturnSendMessage struct {
	Ok          bool              `json:"ok" bson:"ok" query:"ok" form:"ok"`
	Result      ResultSendMessage `json:"result" bson:"result" query:"result" form:"result"`
	HTTPCODE    int
	HTTPMessage string
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

// SendMessageReplyMode this function for send message to the chatID with reply the previous Message
func SendMessageReplyMode(chatID string, message string, replyID string, token string) (ReturnSendMessage, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/sendMessage"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)

	q := request.URL.Query()
	q.Add("chat_id", chatID)
	q.Add("text", message)
	q.Add("reply_to_message_id", replyID)

	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSendMessage{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSendMessage
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// SendMessage this function for send message to the requester
func SendMessage(chatID string, message string, token string) (ReturnSendMessage, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/sendMessage"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)

	q := request.URL.Query()
	q.Add("chat_id", chatID)
	q.Add("text", message)

	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSendMessage{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSendMessage
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// SendMessageReplyModeWithParseOption this function for send message to the chatID with reply the previous Message with parse option
func SendMessageReplyModeWithParseOption(chatID string, message string, replyID string, token string, parseOption string) (ReturnSendMessage, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/sendMessage"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)

	q := request.URL.Query()
	q.Add("chat_id", chatID)
	q.Add("text", message)
	q.Add("reply_to_message_id", replyID)
	q.Add("parse_mode", parseOption)

	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSendMessage{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSendMessage
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// SendMessageWithParseOption this function for send message to the requester With Parse Option
func SendMessageWithParseOption(chatID string, message string, token string, parseOption string) (ReturnSendMessage, error) {
	telegramURLOfficial := "https://api.telegram.org/bot" + token + "/sendMessage"

	request, _ := http.NewRequest("POST", telegramURLOfficial, nil)

	q := request.URL.Query()
	q.Add("chat_id", chatID)
	q.Add("text", message)
	q.Add("parse_mode", parseOption)

	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return ReturnSendMessage{}, err
	}
	defer response.Body.Close()

	// make the data more human vision
	var returnAction ReturnSendMessage
	returnAction.HTTPCODE = response.StatusCode
	returnAction.HTTPMessage = response.Status
	json.NewDecoder(response.Body).Decode(&returnAction)

	return returnAction, nil
}

// Comming Soon
// SetWebhookWithCertificate this function for setting webbook into telegram official API with Certificate
