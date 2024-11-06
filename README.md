# Courier CLI

### Build, test, and manage your [Courier](https://www.courier.com) integration directly from the command line.

The Courier CLI is a developer tool to help you build, test, and manage your integration with Courier directly from the command line. It’s simple to install, works on macOS, Windows, and Linux, and offers a range of functionality to enhance your developer experience with Courier. You can use the Courier CLI to:

- Send messages from the command line to users, lists, or audiences
- Track events to trigger your Courier Automations
- Push & pull industry-standard translation files for internationalizing your content

## Installing the Courier CLI

From the command-line, use the following command to install the Courier CLI and set your API key:

```bash
$ yarn install -g @trycourier/cli
$ courier config --apikey <your-api-key>
```

### Requirements

- Courier CLI has only been tested on node.js v18+

## Authenticate the CLI

The fastest way to get started is to run:

```
$ courier config --apikey <your-api-key>
```

Courier CLI looks for environment variables prefixed with `COURIER_AUTH_TOKEN`. It will load keys from the first location it finds in the following list:

- A `.courier` file in the current working directory
- `~/.courier` (in your home directory)
- A `COURIER_AUTH_TOKEN` or `COURIER_AUTH_TOKEN_*` value otherwise set in your environment (such as via `~/.profile` or `~/.zshrc`)

You can find your Courier API key in your [Courier Settings](https://app.courier.com/settings/api-keys).

## Commands

- `courier config` – Set your Courier API key
- `courier whoami` – Display the currently authenticated workspace
- `courier send` - Send a notification to a user, list, tenant, or audience
- `courier track` - Send a track event to trigger a Courier Automations
- `courier users:get` - Fetch the data for a given user ID
- `courier users:set` - Overwrite a user's profile with the provided data
- `courier users:bulk` - Bulk upload users via csv, json, or parquet
- `courier translations:upload` - Upload .PO files to your Courier workspace
- `courier translations:download` - Download .PO files from your Courier workspace

For more details, run `courier` to see a list of commands and their arguments & options.

## Examples

```
courier --help
courier --version
courier upgrade

courier send --tel 555-867-5309 --body "Hey Jenny\!"
courier send --user user123 --template my-template-id --foo bar
courier send -P --user=test123 --body "hello world" --title="hello" --channels=inbox
courier send --tenant=kewl --title=hello --body="hello world" --channel=inbox
courier send --user="1" --tenant-context=kewl --title=hello --body="hello world" --channel=inbox

courier users:get user123
courier users:set user123 --email user@example.com
courier users:bulk examples/users.csv --replace
courier users:bulk examples/users.parquet --list new-list-id --tenant new-tenant-id

courier track EXAMPLE_EVENT user123 --name "Pip the Pigeon"

courier translations:upload en-US ./translations/en-US.po
courier translations:download en-US --text > example.en-US.po

courier config --apikey MY_API_KEY -P --override
courier config --apikey MY_API_KEY --mock
courier config --apikey MY_API_KEY --draft

courier test-user123 --scopes=read:user-tokens,write:user-tokens --expiration=60
courier test-user123 --all --quiet | pbcopy
```

## Common Flags

There are a number flags you can use for any command

| Flags                      | Description                                                                                                              |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| -M --mock                  | Use the API key that simulates sending using the simulating routing                                                      |
| -P --production            | Use the production environment API key                                                                                   |
| -D --draft                 | Use the draft document scope API key. Use draft or submitted, will default to published key if neither are provided      |
| -S --submitted             | Use the submitted document scope API key                                                                                 |
| --apikey <Courier API Key> | Use the provided Courier API key, otherwise use the approprate environment variable                                      |
| --apiurl <Courier API URL> | Use the provided Courier API URL, otherwise use COURIER_API_URL environment variable. Default is https://api.courier.com |

## Misc

- If you need to change the Courier API URL, you can set COURIER_API_URL in .courier or other methods to set the environment variables.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[Courier](https://github.com/trycourier) ([support@courier.com](mailto:support@courier.com))
