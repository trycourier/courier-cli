# Courier CLI

The Courier CLI sends notifications, inspects delivery, and manages Courier resources from your terminal or a script. It covers the same API surface as the SDKs, so anything you can send from code you can send from the command line.

## Installation

```bash
npm install -g @trycourier/cli
```

This downloads a platform-specific native binary in a postinstall step — no Node.js runtime is needed afterwards.

With Go instead (requires Go 1.22+):

```bash
go install 'github.com/trycourier/courier-cli/v3/cmd/courier@latest'
```

That puts the binary in `$(go env GOPATH)/bin`, which needs to be on your `PATH`.

## Quick Start

```bash
export COURIER_API_KEY='your-api-key'

courier send message \
  --message '{"to": {"user_id": "your_user_id"}, "template": "your_template_id"}'
```

Every command follows the same shape:

```bash
courier [resource] <command> [flags...]
```

Use `--help` on any command to see its flags, and `--api-key` to override the `COURIER_API_KEY` environment variable.

## Documentation

Full documentation: **[courier.com/docs/tools/cli](https://www.courier.com/docs/tools/cli/)**

- [Quickstart](https://www.courier.com/docs/getting-started/quickstart/)
- [Send API](https://www.courier.com/docs/platform/sending/send-message/)
- [API Reference](https://www.courier.com/docs/reference/get-started/)
