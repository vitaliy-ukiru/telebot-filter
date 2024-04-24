// Package dispatcher provides full routing mechanisms.
// Includes filters and upgraded middlewares.
//
// It abstraction on telebot Bot, and it may break
// handlers without dispatcher.
//
// You can use addons, that support this module or
// are suitable for a standard telebot interface.
//
// # Middlewares
//
// With this package you can create complex system with routing and middlewares.
// You can inject middlewares at:
//   - telebot layer (Use method)
//   - dispatcher endpoints (tele.OnText, tele.OnVideo, etc...)
//   - to every router (sub routes will execute parent's middlewares recursive)
//   - to separate handler
//
// Building chain of middlewares will after filters check.
// If you want to execute middleware before filters you need
// setup middleware on endpoint (dispatcher.UseOn) or
// telebot global middlewares.
//
// # Handlers
//
// Exists 3 base methods for add middlewares:
//   - [Router.Bind]
//   - [Router.Handle]
//   - [Router.Dispatch]
//
// or their wrappers from the [Dispatcher] object with same names.
//
// You can read details of these methods by links to documentation.
package dispatcher
