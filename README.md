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
    telefilter.NewRawHandler(
        "/start",
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
No need setup.

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

# Structure

Module separated on two packages: dispatcher and telefilter.

Let's go deep.

## dispatcher package

This package provides Dispatcher, Router and Builder types.
They are responsible for routing new updates.

### Dispatcher & router

The dispatcher execute chain of handling event.
But for register handler uses Router.
Dispatcher stores inside main router. And you can create children routers.
With this you can inherit the middleware.

Also, dispatcher can set middlewares to endpoint.
We don't do this like in default telebot.

### Handler

Handler object is checks and processing updates.

Package provided builtin implementation - `dispatcher.RawHandler`.
For filtering it uses slice of `dispatcher.Filter`.
All filters must return a positive result for further work.
This behavior is not described explicitly anywhere, but is assumed.

### Route

This is a service structure.
It combines a handler and a list of middleware.
Required for sharing responsibilities.
The handler should not talk about middleware.
It may help to reuse handlers

This division can also help when creating custom add-ons.
Providing an interaction interface convenient for other people's code

### Chain of execute of handler

```
# The lower the level, the closer to the handler
1. Bot scoped middlewares
2. Dispatcher's group middlewares 
3. Dispatcher's endpoint middlewares.
4. Middlewares of router chain.
5. Handler's Check method   # Method for check update on match
6. Handler's Execute method # Method of process update.
```

Builder is helper for creating handlers

## The routing package

This package is more simplify implementation for filter support.
Package is full back compatibility with standard telebot code.

But this package don't provide middlewares as in the dispatcher package.
It just gives generic handler.

Another catch is registering handlers for one endpoint
must pass in one call to `bot.Handle`

```go
bot.Handle("/start", routing.New(
    ... // your handlers
))

// in another place down the code
bot.Handle("/start", routing.New( // !! Only this handler will be registered
    anotherHandler,
))

```