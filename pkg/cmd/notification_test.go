// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
)

func TestNotificationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "create",
			"--notification", "{brand: {id: bnd_01kx4mrd0pfzw8wt7pn7p2fzag}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, title: Welcome!, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: 'Hello {{data.name}}.', font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}], version: '2022-01-01'}, name: Welcome Email, routing: {strategy_id: rs_01kx4h2jdafq8bk9amzvy6hbv0}, subscription: {topic_id: pt_01kx4h2jdafq8bk9a26x0kvd1t}, tags: [onboarding, welcome], alias: welcome}",
			"--state", "DRAFT",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notification:\n" +
			"  brand:\n" +
			"    id: bnd_01kx4mrd0pfzw8wt7pn7p2fzag\n" +
			"  content:\n" +
			"    elements:\n" +
			"      - channels:\n" +
			"          - string\n" +
			"        if: if\n" +
			"        loop: loop\n" +
			"        ref: ref\n" +
			"        channel: email\n" +
			"        elements:\n" +
			"          - channels:\n" +
			"              - string\n" +
			"            if: if\n" +
			"            loop: loop\n" +
			"            ref: ref\n" +
			"            title: Welcome!\n" +
			"            type: meta\n" +
			"          - channels:\n" +
			"              - string\n" +
			"            if: if\n" +
			"            loop: loop\n" +
			"            ref: ref\n" +
			"            align: left\n" +
			"            bold: bold\n" +
			"            color: color\n" +
			"            content: Hello {{data.name}}.\n" +
			"            font_size: font_size\n" +
			"            format: markdown\n" +
			"            italic: italic\n" +
			"            line_height: line_height\n" +
			"            locales:\n" +
			"              foo:\n" +
			"                content: content\n" +
			"            strikethrough: strikethrough\n" +
			"            text_style: text\n" +
			"            underline: underline\n" +
			"            type: text\n" +
			"        font_size: font_size\n" +
			"        line_height: line_height\n" +
			"        padding: padding\n" +
			"        raw:\n" +
			"          foo: bar\n" +
			"        type: channel\n" +
			"    version: '2022-01-01'\n" +
			"  name: Welcome Email\n" +
			"  routing:\n" +
			"    strategy_id: rs_01kx4h2jdafq8bk9amzvy6hbv0\n" +
			"  subscription:\n" +
			"    topic_id: pt_01kx4h2jdafq8bk9a26x0kvd1t\n" +
			"  tags:\n" +
			"    - onboarding\n" +
			"    - welcome\n" +
			"  alias: welcome\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "create",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestNotificationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "retrieve",
			"--id", "id",
			"--version", "version",
		)
	})
}

func TestNotificationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "list",
			"--cursor", "cursor",
			"--event-id", "event_id",
			"--notes=true",
		)
	})
}

func TestNotificationsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "archive",
			"--id", "id",
		)
	})
}

func TestNotificationsGetMetrics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "get-metrics",
			"--id", "x",
			"--end", "'2019-12-27T18:11:19.117Z'",
			"--granularity", "HOUR",
			"--lookback", "lookback",
			"--start", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestNotificationsListVersions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "list-versions",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNotificationsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "publish",
			"--id", "id",
			"--version", "v321669910225",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: v321669910225")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "publish",
			"--id", "id",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestNotificationsPutContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "put-content",
			"--id", "id",
			"--content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, title: Welcome!, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: 'Hello {{data.name}}.', font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}], version: '2022-01-01'}",
			"--state", "DRAFT",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsPutContent)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "put-content",
			"--id", "id",
			"--content.elements", "[{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, title: Welcome!, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: 'Hello {{data.name}}.', font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}]",
			"--content.version", "2022-01-01",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content:\n" +
			"  elements:\n" +
			"    - channels:\n" +
			"        - string\n" +
			"      if: if\n" +
			"      loop: loop\n" +
			"      ref: ref\n" +
			"      channel: email\n" +
			"      elements:\n" +
			"        - channels:\n" +
			"            - string\n" +
			"          if: if\n" +
			"          loop: loop\n" +
			"          ref: ref\n" +
			"          title: Welcome!\n" +
			"          type: meta\n" +
			"        - channels:\n" +
			"            - string\n" +
			"          if: if\n" +
			"          loop: loop\n" +
			"          ref: ref\n" +
			"          align: left\n" +
			"          bold: bold\n" +
			"          color: color\n" +
			"          content: Hello {{data.name}}.\n" +
			"          font_size: font_size\n" +
			"          format: markdown\n" +
			"          italic: italic\n" +
			"          line_height: line_height\n" +
			"          locales:\n" +
			"            foo:\n" +
			"              content: content\n" +
			"          strikethrough: strikethrough\n" +
			"          text_style: text\n" +
			"          underline: underline\n" +
			"          type: text\n" +
			"      font_size: font_size\n" +
			"      line_height: line_height\n" +
			"      padding: padding\n" +
			"      raw:\n" +
			"        foo: bar\n" +
			"      type: channel\n" +
			"  version: '2022-01-01'\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "put-content",
			"--id", "id",
		)
	})
}

func TestNotificationsPutElement(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "put-element",
			"--id", "id",
			"--element-id", "elementId",
			"--type", "text",
			"--channel", "string",
			"--data", "{content: bar}",
			"--if", "if",
			"--loop", "loop",
			"--ref", "ref",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: text\n" +
			"channels:\n" +
			"  - string\n" +
			"data:\n" +
			"  content: bar\n" +
			"if: if\n" +
			"loop: loop\n" +
			"ref: ref\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "put-element",
			"--id", "id",
			"--element-id", "elementId",
		)
	})
}

func TestNotificationsPutLocale(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "put-locale",
			"--id", "id",
			"--locale-id", "localeId",
			"--element", "{id: elem_1}",
			"--element", "{id: elem_2}",
			"--state", "DRAFT",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsPutLocale)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "put-locale",
			"--id", "id",
			"--locale-id", "localeId",
			"--element.id", "elem_1",
			"--element.id", "elem_2",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"elements:\n" +
			"  - id: elem_1\n" +
			"  - id: elem_2\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "put-locale",
			"--id", "id",
			"--locale-id", "localeId",
		)
	})
}

func TestNotificationsReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "replace",
			"--id", "id",
			"--notification", "{brand: {id: id}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, title: Updated, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: Updated content., font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}], version: '2022-01-01'}, name: Updated Name, routing: {strategy_id: strategy_id}, subscription: {topic_id: topic_id}, tags: [updated], alias: alias}",
			"--state", "PUBLISHED",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notification:\n" +
			"  brand:\n" +
			"    id: id\n" +
			"  content:\n" +
			"    elements:\n" +
			"      - channels:\n" +
			"          - string\n" +
			"        if: if\n" +
			"        loop: loop\n" +
			"        ref: ref\n" +
			"        channel: email\n" +
			"        elements:\n" +
			"          - channels:\n" +
			"              - string\n" +
			"            if: if\n" +
			"            loop: loop\n" +
			"            ref: ref\n" +
			"            title: Updated\n" +
			"            type: meta\n" +
			"          - channels:\n" +
			"              - string\n" +
			"            if: if\n" +
			"            loop: loop\n" +
			"            ref: ref\n" +
			"            align: left\n" +
			"            bold: bold\n" +
			"            color: color\n" +
			"            content: Updated content.\n" +
			"            font_size: font_size\n" +
			"            format: markdown\n" +
			"            italic: italic\n" +
			"            line_height: line_height\n" +
			"            locales:\n" +
			"              foo:\n" +
			"                content: content\n" +
			"            strikethrough: strikethrough\n" +
			"            text_style: text\n" +
			"            underline: underline\n" +
			"            type: text\n" +
			"        font_size: font_size\n" +
			"        line_height: line_height\n" +
			"        padding: padding\n" +
			"        raw:\n" +
			"          foo: bar\n" +
			"        type: channel\n" +
			"    version: '2022-01-01'\n" +
			"  name: Updated Name\n" +
			"  routing:\n" +
			"    strategy_id: strategy_id\n" +
			"  subscription:\n" +
			"    topic_id: topic_id\n" +
			"  tags:\n" +
			"    - updated\n" +
			"  alias: alias\n" +
			"state: PUBLISHED\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "replace",
			"--id", "id",
		)
	})
}

func TestNotificationsRetrieveContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "retrieve-content",
			"--id", "id",
			"--version", "version",
		)
	})
}
