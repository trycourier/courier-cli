// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestBrandsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "create",
			"--api-key", "string",
			"--name", "name",
			"--id", "id",
			"--settings", "{colors: {primary: primary, secondary: secondary}, email: {footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}, inapp: {colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}}",
			"--snippets", "{items: [{name: name, value: value}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(brandsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "create",
			"--api-key", "string",
			"--name", "name",
			"--id", "id",
			"--settings.colors", "{primary: primary, secondary: secondary}",
			"--settings.email", "{footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}",
			"--settings.inapp", "{colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}",
			"--snippets.items", "[{name: name, value: value}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"id: id\n" +
			"settings:\n" +
			"  colors:\n" +
			"    primary: primary\n" +
			"    secondary: secondary\n" +
			"  email:\n" +
			"    footer:\n" +
			"      content: content\n" +
			"      inheritDefault: true\n" +
			"    head:\n" +
			"      inheritDefault: true\n" +
			"      content: content\n" +
			"    header:\n" +
			"      logo:\n" +
			"        href: href\n" +
			"        image: image\n" +
			"      barColor: barColor\n" +
			"      inheritDefault: true\n" +
			"    templateOverride:\n" +
			"      enabled: true\n" +
			"      backgroundColor: backgroundColor\n" +
			"      blocksBackgroundColor: blocksBackgroundColor\n" +
			"      footer: footer\n" +
			"      head: head\n" +
			"      header: header\n" +
			"      width: width\n" +
			"      mjml:\n" +
			"        enabled: true\n" +
			"        backgroundColor: backgroundColor\n" +
			"        blocksBackgroundColor: blocksBackgroundColor\n" +
			"        footer: footer\n" +
			"        head: head\n" +
			"        header: header\n" +
			"        width: width\n" +
			"      footerBackgroundColor: footerBackgroundColor\n" +
			"      footerFullWidth: true\n" +
			"  inapp:\n" +
			"    colors:\n" +
			"      primary: primary\n" +
			"      secondary: secondary\n" +
			"    icons:\n" +
			"      bell: bell\n" +
			"      message: message\n" +
			"    widgetBackground:\n" +
			"      bottomColor: bottomColor\n" +
			"      topColor: topColor\n" +
			"    borderRadius: borderRadius\n" +
			"    disableMessageIcon: true\n" +
			"    fontFamily: fontFamily\n" +
			"    placement: top\n" +
			"snippets:\n" +
			"  items:\n" +
			"    - name: name\n" +
			"      value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "brands", "create",
			"--api-key", "string",
		)
	})
}

func TestBrandsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "retrieve",
			"--api-key", "string",
			"--brand-id", "brand_id",
		)
	})
}

func TestBrandsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "update",
			"--api-key", "string",
			"--brand-id", "brand_id",
			"--name", "name",
			"--settings", "{colors: {primary: primary, secondary: secondary}, email: {footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}, inapp: {colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}}",
			"--snippets", "{items: [{name: name, value: value}]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(brandsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "update",
			"--api-key", "string",
			"--brand-id", "brand_id",
			"--name", "name",
			"--settings.colors", "{primary: primary, secondary: secondary}",
			"--settings.email", "{footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}",
			"--settings.inapp", "{colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}",
			"--snippets.items", "[{name: name, value: value}]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"settings:\n" +
			"  colors:\n" +
			"    primary: primary\n" +
			"    secondary: secondary\n" +
			"  email:\n" +
			"    footer:\n" +
			"      content: content\n" +
			"      inheritDefault: true\n" +
			"    head:\n" +
			"      inheritDefault: true\n" +
			"      content: content\n" +
			"    header:\n" +
			"      logo:\n" +
			"        href: href\n" +
			"        image: image\n" +
			"      barColor: barColor\n" +
			"      inheritDefault: true\n" +
			"    templateOverride:\n" +
			"      enabled: true\n" +
			"      backgroundColor: backgroundColor\n" +
			"      blocksBackgroundColor: blocksBackgroundColor\n" +
			"      footer: footer\n" +
			"      head: head\n" +
			"      header: header\n" +
			"      width: width\n" +
			"      mjml:\n" +
			"        enabled: true\n" +
			"        backgroundColor: backgroundColor\n" +
			"        blocksBackgroundColor: blocksBackgroundColor\n" +
			"        footer: footer\n" +
			"        head: head\n" +
			"        header: header\n" +
			"        width: width\n" +
			"      footerBackgroundColor: footerBackgroundColor\n" +
			"      footerFullWidth: true\n" +
			"  inapp:\n" +
			"    colors:\n" +
			"      primary: primary\n" +
			"      secondary: secondary\n" +
			"    icons:\n" +
			"      bell: bell\n" +
			"      message: message\n" +
			"    widgetBackground:\n" +
			"      bottomColor: bottomColor\n" +
			"      topColor: topColor\n" +
			"    borderRadius: borderRadius\n" +
			"    disableMessageIcon: true\n" +
			"    fontFamily: fontFamily\n" +
			"    placement: top\n" +
			"snippets:\n" +
			"  items:\n" +
			"    - name: name\n" +
			"      value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "brands", "update",
			"--api-key", "string",
			"--brand-id", "brand_id",
		)
	})
}

func TestBrandsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "list",
			"--api-key", "string",
			"--cursor", "cursor",
		)
	})
}

func TestBrandsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "brands", "delete",
			"--api-key", "string",
			"--brand-id", "brand_id",
		)
	})
}
