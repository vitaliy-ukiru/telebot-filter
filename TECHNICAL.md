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