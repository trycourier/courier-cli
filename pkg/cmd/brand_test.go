// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestBrandsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "create",
		"--api-key", "string",
		"--name", "name",
		"--id", "id",
		"--settings", "{colors: {primary: primary, secondary: secondary}, email: {footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}, inapp: {colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}}",
		"--snippets", "{items: [{name: name, value: value}]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(brandsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "create",
		"--name", "name",
		"--id", "id",
		"--settings.colors", "{primary: primary, secondary: secondary}",
		"--settings.email", "{footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}",
		"--settings.inapp", "{colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}",
		"--snippets.items", "[{name: name, value: value}]",
	)
}

func TestBrandsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "retrieve",
		"--api-key", "string",
		"--brand-id", "brand_id",
	)
}

func TestBrandsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "update",
		"--api-key", "string",
		"--brand-id", "brand_id",
		"--name", "name",
		"--settings", "{colors: {primary: primary, secondary: secondary}, email: {footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}, inapp: {colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}}",
		"--snippets", "{items: [{name: name, value: value}]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(brandsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "update",
		"--brand-id", "brand_id",
		"--name", "name",
		"--settings.colors", "{primary: primary, secondary: secondary}",
		"--settings.email", "{footer: {content: content, inheritDefault: true}, head: {inheritDefault: true, content: content}, header: {logo: {href: href, image: image}, barColor: barColor, inheritDefault: true}, templateOverride: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width, mjml: {enabled: true, backgroundColor: backgroundColor, blocksBackgroundColor: blocksBackgroundColor, footer: footer, head: head, header: header, width: width}, footerBackgroundColor: footerBackgroundColor, footerFullWidth: true}}",
		"--settings.inapp", "{colors: {primary: primary, secondary: secondary}, icons: {bell: bell, message: message}, widgetBackground: {bottomColor: bottomColor, topColor: topColor}, borderRadius: borderRadius, disableMessageIcon: true, fontFamily: fontFamily, placement: top}",
		"--snippets.items", "[{name: name, value: value}]",
	)
}

func TestBrandsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "list",
		"--api-key", "string",
		"--cursor", "cursor",
	)
}

func TestBrandsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "delete",
		"--api-key", "string",
		"--brand-id", "brand_id",
	)
}
