# Courier CLI

### Build, test, and manage your [Courier](https://www.courier.com) integration directly from the command line.

The Courier CLI is a developer tool to help you build, test, and manage your integration with Courier directly from the command line. It’s simple to install, works on macOS, Windows, and Linux, and offers a range of functionality to enhance your developer experience with Courier. You can use the Courier CLI to:

- Send messages from the command line to users, lists, or audiences
- Track events to trigger your Courier Automations
- Push & pull industry-standard translation files for internationalizing your content

## Installing the Courier CLI

From the command-line, use the following command to install the Courier CLI and set your API key:

```bash
$ npm install -g @trycourier/cli
$ courier config --apikey <your-api-key>
```

### Requirements

- Courier CLI has only been tested on node.js v16+

## Authenticate the CLI

The fastest way to get started is to run:

```
$ courier config --apikey <your-api-key>
```

Courier CLI looks for an environment variable named `COURIER_API_KEY`. It will load that key from the first location it finds in the following list:

- A `.courier` file in the current working directory
- `~/.courier` (in your home directory)
- A `COURIER_API_KEY` value otherwise set in your environment (such as via `~/.profile` or `~/.zshrc`)

You can find your Courier API key in your [Courier Settings](https://app.courier.com/settings/api-keys).

## Commands

- `courier config` – Set your Courier API key in `~/.courier`
- `courier whoami` – Display the currently authenticated workspace
- `courier send` - Send a notification to a user, list, or audience
- `courier track` - Send an event to test your Courier Automations
- `courier digests:flush` - Send an event to test your Courier Automations
- `courier translations:upload` - Upload .PO files to your Courier workspace
- `courier translations:download` - Download .PO files from your Courier workspace

For more details, run `courier` to see a list of commands and their arguments & options.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[Courier](https://github.com/trycourier) ([support@courier.com](mailto:support@courier.com))
