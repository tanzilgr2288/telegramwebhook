package telegramwebhook_test

import (
	"net/http"
	"telegramwebhook"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// @author Kevin Setiawan Tanzil [Tom Cruise]
// @desc this file for handler simple webhook bind & unbind to telegram API Testing file
// @file telegramWebhook_test.go
// @since 29 Maret 2019

// initialitation and Declaration variable
var (
	callbackURL = "callback_url"
	token       = "token"
)

// TestSetWebhook with mock data
func TestSetWebhook(t *testing.T) {
	allowUpdateMock := make([]string, 1)
	allowUpdateMock[0] = "message"

	testingResult, _ := telegramwebhook.SetWebhook(callbackURL, 100, allowUpdateMock, token)

	assert.Equal(t, testingResult.HTTPCODE, http.StatusOK)

}

// TestSetWebhook with mock data and max connection -1 so it will be become 40
func TestSetWebhookMaxConnectionLessThanOne(t *testing.T) {
	time.Sleep(2 * time.Second) // for delaying the test, so telegram will not reject
	allowUpdateMock := make([]string, 1)
	allowUpdateMock[0] = "message"

	testingResult, _ := telegramwebhook.SetWebhook(callbackURL, -1, allowUpdateMock, token)

	assert.Equal(t, testingResult.HTTPCODE, http.StatusOK)

}

// TestGetWebhookInfo this is for testing getWebhook Information
func TestGetWebhookInfo(t *testing.T) {
	time.Sleep(2 * time.Second) // for delaying the test, so telegram will not reject
	allowUpdateMock := make([]string, 1)
	allowUpdateMock[0] = "message"

	testingResult, _ := telegramwebhook.GetWebhookInfo(token)

	assert.Equal(t, testingResult.HTTPCODE, http.StatusOK)
	assert.Equal(t, testingResult.Result.URL, callbackURL)

}

// TestDeleteWebHookInstance this is for testing delete webhook instance
func TestDeleteWebHookInstance(t *testing.T) {
	time.Sleep(2 * time.Second) // for delaying the test, so telegram will not reject
	allowUpdateMock := make([]string, 1)
	allowUpdateMock[0] = "message"

	testingResult, _ := telegramwebhook.DeleteWebHookInstance(token)

	assert.Equal(t, testingResult.HTTPCODE, http.StatusOK)
	assert.Equal(t, testingResult.OK, true)
	assert.Equal(t, testingResult.Result, true)

}
