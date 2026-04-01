# Courier CLI

Courier is a notifications API for sending messages across email, SMS, push, in-app inbox, Slack, and WhatsApp from a single API call. The CLI exposes every API endpoint as a shell command with structured JSON output.

## Setup

```sh
npm install -g @trycourier/cli
export COURIER_API_KEY=your_api_key
```

## Core pattern

```sh
courier send message \
  --message.to.user_id "user_123" \
  --message.template "TEMPLATE_ID" \
  --message.data '{"order_id": "456"}' \
  --message.routing.method "single" \
  --message.routing.channels '["email", "sms"]'
```

## Key rules

- Use `--message.routing.method "single"` (fallback chain) unless the user explicitly asks for parallel delivery (`"all"`).
- Use `courier profiles create` for partial profile updates (it merges). Use `courier profiles replace` only when fully replacing all profile data.
- Test and production use different API keys from the same workspace. Always confirm which environment with `--debug`.
- Bulk sends are a 3-step flow: `courier bulk create-job` → `courier bulk add-users` → `courier bulk run-job`.
- Pass `--format json` for machine-readable output; default `auto` format is human-friendly.
- Use `--help` on any command to see all available flags.

## Common commands

```sh
courier send message --message '{...}'       # send a notification
courier messages list --recipient "user_123" # check delivery status
courier profiles create "user_123" --profile '{"email":"a@b.com"}'
courier profiles replace "user_123" --profile '{"email":"a@b.com"}'
courier bulk create-job --message '{...}'    # start bulk send
courier bulk add-users JOB_ID --users '[...]'
courier bulk run-job JOB_ID
```

## Concepts

- `template` — notification template ID from the Courier dashboard
- `routing.method` — `"single"` = try channels in order until one succeeds; `"all"` = send on every channel simultaneously
- `tenant_id` — multi-tenant context; affects brand and preference defaults for the message
- `list_id` — send to all subscribers of a named list

## More context

- Full docs index: https://www.courier.com/docs/llms.txt
- API reference: https://www.courier.com/docs/reference/get-started
- MCP server (richer IDE integration): https://mcp.courier.com
- Courier Skills (Cursor / Claude Code): https://github.com/trycourier/courier-skills
