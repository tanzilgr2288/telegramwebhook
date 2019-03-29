# Golang Simple Telegram Webhook Binding Library

This Library can help you to not using the whole overkill telegram Library,
it can save your time to develope telegram webhook just for binding the callback URL

## What have done
1. Set WebHook Without Cert
2. Get WebHook Info
3. Unbind / delete Webhook instance from telegram

## What should do in the future
1. Set Webhook with cert

## Coverage Test : 91%

# How to run test
please go to the telegramWebhook_test.go and modif the callbackUrl and Token
1. For callback URL testing using [Webhook Mocking URL](https://webhook.site)
2. For Token, use your own token

# How to use
First, ensure the library is installed and up to date by running
`go get -u github.com/tanzilgr2288/telegramwebhook`.

## Set up the webhook 

### The Parameter Explanation
1. callbackURL -> This is your callback URL When webhook send message to you when your bot receiving message
2. MaximumConnection -> Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. range less than 1 and more than 100, will be converted to default value 40
3. allowUpdate -> List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. list of type can see on [Telegram Allow Update List](https://core.telegram.org/bots/api#update)
4. Token -> your bot Token, it receive from BotFather [Create Telegram Bot](https://core.telegram.org/bots#botfather)
```go
package main

import (
	"fmt"
	"github.com/tanzilgr2288/telegramwebhook"
)

func main() {
  // SetWebhook(callbackURL string, maxConnection int, allowUpdate []string, token string)
  // will return you ReturnSetWebHookAndDelete struct
	result, err := telegramwebhook.SetWebhook("your_callback_url", "maximum_connection", "your_allow_update_list", "token")
}

```
## Get Webhook Info
```go
package main

import (
	"fmt"
	"github.com/tanzilgr2288/telegramwebhook"
)

func main() {
  // GetWebhookInfo(token string)
  // will return you ReturnGetWebHookInfo struct
	result, err := telegramwebhook.GetWebhookInfo("token")
}

```

## Delete Webhook Instance
```go
package main

import (
	"fmt"
	"github.com/tanzilgr2288/telegramwebhook"
)

func main() {
  // DeleteWebHookInstance(token string)
  // will return you ReturnSetWebHookAndDelete struct
	result, err := telegramwebhook.DeleteWebHookInstance("token")
}

```
