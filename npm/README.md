# Courier CLI

The official CLI for the [Courier REST API](https://www.courier.com/docs). Manage notifications, templates, users, and more from the command line.

## Installation

```sh
npm install -g @trycourier/cli
```

### Other install methods

- **Binary**: [GitHub Releases](https://github.com/trycourier/courier-cli/releases)

## Usage

```sh
export COURIER_API_KEY=your_api_key

courier send message --message.to.email "user@example.com" --message.template "my-template"
courier messages list
courier profiles retrieve --user-id "user-123"
```

For details about specific commands, use `courier --help`.

## How it works

This package downloads the platform-specific Courier CLI binary from [GitHub Releases](https://github.com/trycourier/courier-cli/releases) during `npm install`. No Node.js runtime dependency at execution time.

## Documentation

- [Courier Docs](https://www.courier.com/docs)
- [CLI Repository](https://github.com/trycourier/courier-cli)
- [API Reference](https://www.courier.com/docs/reference)
