# disgohook

[![Go Reference](https://pkg.go.dev/badge/github.com/DisgoOrg/disgohook.svg)](https://pkg.go.dev/github.com/DisgoOrg/disgohook)
[![Go Report](https://goreportcard.com/badge/github.com/DisgoOrg/disgohook)](https://goreportcard.com/report/github.com/DisgoOrg/disgohook)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DisgoOrg/disgohook)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/disgohook/blob/master/LICENSE)
[![Disgo Version](https://img.shields.io/github/v/release/DisgoOrg/disgohook)](https://github.com/DisgoOrg/disgohook/releases/latest)
[![Disgo Discord](https://img.shields.io/badge/Disgo%20Discord-blue.svg)](https://discord.gg/mgjJeufk)

DisgoHook is a simple [Discord Webhook](https://discord.com/developers/docs/resources/webhook) library written
in [Go](https://golang.org/) aimed for simplicity and easy use

## Getting Started

### Installing

```sh
go get github.com/DisgoOrg/disgohook
```

### Usage

Import the package into your project.

```go
import "github.com/DisgoOrg/disgohook"
```

Create a new Webhook by `webhook_id/webhook_token` and pass a [logger](https://github.com/DisgoOrg/log)
like [logrus](https://github.com/sirupsen/logrus). This webhook then can be used to send/edit/delete messages

```go
logger := logrus.New()
webhook, err := disgohook.NewWebhookByToken(nil, logger, "webhook_id/webhook_token")

message, err := webhook.SendContent("hello world!")
```

## Documentation

Documentation is unfinished and can be found under

* [![Go Reference](https://pkg.go.dev/badge/github.com/DisgoOrg/disgohook.svg)](https://pkg.go.dev/github.com/DisgoOrg/disgohook)
* [![Discord Webhook Documentation](https://img.shields.io/badge/Discord%20Webhook%20Documentation-blue.svg)](https://discord.com/developers/docs/resources/webhook)

## Examples

You can find examples under [example](https://github.com/DisgoOrg/disgohook/tree/master/example)
and [DisLog](https://github.com/DisgoOrg/dislog)

## Troubleshooting

For help feel free to open an issues or reach out on [Discord](https://discord.gg/mgjJeufk)

## Contributing

Contributions are welcomed but for bigger changes please first reach out via [Discord](https://discord.gg/mgjJeufk) or
create an issue to discuss your intetions and ideas.

## License

Distributed under
the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/disgohook/blob/master/LICENSE)
. See LICENSE for more information.
