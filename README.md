# telebot-filter

# Motivation

This module provides use telebot module but with filters
and upgraded middlewares system.

The module is also compatible with standard handlers.
Provided that their end points do not intersect with
endpoints of this module.

# Quick Start

## Install module

```
go get github.com/vitaliy-ukiru/telebot-filter
```

## Select flow

You can use different method to register handlers.
These methods separated into packages

- dispatcher
- routing

You can use both package in one program, but endpoints
must not intersect.
The packages do not monitor this behavior.

Read more about their differences and features [here](#structure).

I'll show API of both packages.

### Package dispatcher
#### Make base setup

```go

package main

import (
	"os"

	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tb "gopkg.in/telebot.v3"
)

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("BOT_TOKEN"),
	})
	if err != nil {
		panic(err)
	}

	dp := dispatcher.NewDispatcher(
		// creating group, because the dispatcher
		// does not need access to the bot.
		bot.Group(),
	)
}
```

#### Setups handlers

<details>
<summary>Most likely telebot method</summary>

```go
dp.Handle(
    "/start",
    telefilter.NewRawHandler(   
        func(c tb.Context) error {
            return c.Send("Hi")
        },
        // filters list
        func(c tb.Context) bool {
            return c.Message().Chat.Type == tb.ChatPrivate
        },
    ),
    /* middlewares like in telebot*/
)
```

</details>


<details>
<summary>Using builder</summary>

```go
dp.Bind(
    dp.
    NewHandler(tb.OnText).
    Filter(message.EqualFold("hi")). // from pkg/filters/message,
    Do(func (c tb.Context) error {
        name := c.Message().Sender.FirstName
        return c.Send("Hi, " + name + "!")
    }),
)
```

</details>

<details>
<summary>Using low-level interface</summary>

```go
dp.Dispatch(
    telefilter.NewRoute(
        telefilter.NewRawHandler(
            tb.OnDocument,
            func (c tb.Context) error {
                return c.Send("I'll read this document later")
            },
            // filter
            func (c tb.Context) bool {
                doc := c.Message().Document
                return doc.MIME == "plain/text"
            },
        ),
        // middlewares
        middleware.Whitelist(
            CoolChatID,
        ),
    ),
)
```

Ugly? may be. But you don't use this in normal code.
</details>

#### Add middlewares

You can add middleware to router object, handler and endpoint.

```go
var dp *Dispatcher

dp.UseOn(endpoint, middleware) // middlewares for all handler at endpoint.

dp.Use() // middleware to root router.

router := dp.NewRouter()
router.Use() // middlewares only for this router's handlers.

router.Handle(handler, middleware) // middleware only for handler.

```



### Package routing
Simplest way:
```go
bot.Handle(tele.OnText, routing.New(
	tf.NewHandler(
		handleHi,
		filterHiText,
	),
	tf.NewHandler(
		handleWakeUpInGroup,
		filterGroup,
		filterWakeUpText,
	)
))
```
Alternative. With this method you can add handlers in runtime.
```go

route := new(routing.Route)

route.Add(handlerStartAtFirstTime)
route.Add(handlerStartInGroup)
router.Add(handlerDeeplink)

bot.Handle("/start", route.Handler)
route.Add(handlerStartTwice) // it will added in route

```


#### Setups handlers
```go
bot.Handle(
		"/start",
		routing.New(
			// deeplink handler
			// deeplink is URL t.me/<bot_username>/start=<deeplink>
			// that converted like /start <deeplink>
			// more into at https://core.telegram.org/api/links#bot-links
			tf.NewRawHandler(
				func(c tb.Context) error {
					deeplink := c.Message().Payload
					return c.Send("You deeplink: " + deeplink)
				},

				func(c tb.Context) bool {
					return c.Message().Payload != ""
				},
			),

			// base handler
			tf.NewRawHandler(func(c tb.Context) error {
				return c.Send("Hi!")
			}),
		),

		userDatabaseMiddleware,
	)
```
**Please note that the order in which handlers are registered is important!**
With a different order, we would not have been able to even reach the deeplink filter,
because a basic handler without filters would immediately mark the event as matched.


#### Add middlewares
You can add middlewares only manually like in default telebot.


## Execute bot

Just like in telebot

```go
bot.Start()
```

# Technical information
For more information about internals see [this document](TECHNICAL.md)