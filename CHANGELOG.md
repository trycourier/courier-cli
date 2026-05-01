# Changelog

## 3.5.0 (2026-05-01)

Full Changelog: [v3.4.2...v3.5.0](https://github.com/trycourier/courier-cli/compare/v3.4.2...v3.5.0)

### Features

* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([780a9c0](https://github.com/trycourier/courier-cli/commit/780a9c0e5dc5f889e1756e00b7712b58fd1c2cf6))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([4c52102](https://github.com/trycourier/courier-cli/commit/4c5210271156ddab475383aa6d512d915c43e7b6))
* **cli:** send filename and content type when reading input from files ([8de2857](https://github.com/trycourier/courier-cli/commit/8de28579f6a831bfdc78ec33b16bcf31c55d06bb))
* support passing path and query params over stdin ([29f9b3f](https://github.com/trycourier/courier-cli/commit/29f9b3f993fbd72360b1ce65e77b26e03bbfce8f))


### Bug Fixes

* **cli:** correctly load zsh autocompletion ([c58ffcd](https://github.com/trycourier/courier-cli/commit/c58ffcd91a7c0fb5255b638a34a61c2c8146759c))
* flags for nullable body scalar fields are strictly typed ([41b8096](https://github.com/trycourier/courier-cli/commit/41b809656f1b5a1d4ba73b91e9c60533f032e5f0))


### Chores

* **ci:** support manually triggering release workflow ([99e6955](https://github.com/trycourier/courier-cli/commit/99e6955357951c874d777c4a26a920d602b7ec72))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([09aeebd](https://github.com/trycourier/courier-cli/commit/09aeebd03bb1402bd278940cdeb8ea69b01517b6))
* **cli:** switch long lists of positional args over to param structs ([389c820](https://github.com/trycourier/courier-cli/commit/389c8207c4c54f657dfd310ce5ab1d8d751cc357))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([747c709](https://github.com/trycourier/courier-cli/commit/747c7090e140eb34ff022c5d9e42454618efe5c4))
* **internal:** codegen related update ([b152ea0](https://github.com/trycourier/courier-cli/commit/b152ea0392fc2ee9857dde8affe7746cc07fa0cd))
* **internal:** codegen related update ([cd851ac](https://github.com/trycourier/courier-cli/commit/cd851acc33bb7aa94374078ce84ca0e69064eec3))
* **internal:** codegen related update ([0f66682](https://github.com/trycourier/courier-cli/commit/0f66682369f79620f487884c535e41ad96304d95))
* **internal:** codegen related update ([8f3872f](https://github.com/trycourier/courier-cli/commit/8f3872f856ca43e04183c8f965f8e0e7ec1e0177))
* **internal:** more robust bootstrap script ([3934247](https://github.com/trycourier/courier-cli/commit/39342478242318700e674e0ed42b1950e7bce5a1))


### Documentation

* **api:** update notification parameter descriptions ([40e76a4](https://github.com/trycourier/courier-cli/commit/40e76a403eab7c2c445e532fcbce8494b9cef78a))

## 3.4.2 (2026-04-14)

Full Changelog: [v3.4.1...v3.4.2](https://github.com/trycourier/courier-cli/compare/v3.4.1...v3.4.2)

### Chores

* **internal:** regenerate SDK with no functional changes ([f00bb03](https://github.com/trycourier/courier-cli/commit/f00bb03b74254edfad802f08c495e228f43fc509))

## 3.4.1 (2026-04-13)

Full Changelog: [v3.4.0...v3.4.1](https://github.com/trycourier/courier-cli/compare/v3.4.0...v3.4.1)

### Bug Fixes

* **cli:** fix incompatible Go types for flag generated as array of maps ([213bd1d](https://github.com/trycourier/courier-cli/commit/213bd1d825283670c97ead0231a72fe72071d582))
* fix for failing to drop invalid module replace in link script ([ecb6b7b](https://github.com/trycourier/courier-cli/commit/ecb6b7b4bc868ef0ffede0efd7fd614b277c7282))


### Chores

* **cli:** additional test cases for `ShowJSONIterator` ([7a49de2](https://github.com/trycourier/courier-cli/commit/7a49de24ba244acf937ed42fcdd14b2b38583bd4))


### Documentation

* **api:** update usage text for providers update method ([797aaf8](https://github.com/trycourier/courier-cli/commit/797aaf8b1d7a31cf2981022e2df0bc3ab4441163))

## 3.4.0 (2026-04-09)

Full Changelog: [v3.3.0...v3.4.0](https://github.com/trycourier/courier-cli/compare/v3.3.0...v3.4.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([a902634](https://github.com/trycourier/courier-cli/commit/a902634508d62e7aeaed7f7a32689e7955e38105))
* **api:** add put-content/element/locale to notifications, remove draft resource ([73eb073](https://github.com/trycourier/courier-cli/commit/73eb0735e93404b04352079a847b0bccba564390))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([2efc982](https://github.com/trycourier/courier-cli/commit/2efc9824db1675e76ed76ec684ecbb8535cf6a92))
* binary-only parameters become CLI flags that take filenames only ([a9fe67e](https://github.com/trycourier/courier-cli/commit/a9fe67e4100a9d864ca53a6e3340e8ac6b574c1f))


### Bug Fixes

* fall back to main branch if linking fails in CI ([431a672](https://github.com/trycourier/courier-cli/commit/431a672eb91d5c3fff0c8347d0b2cc38c3aad7ce))
* fix quoting typo ([de9edff](https://github.com/trycourier/courier-cli/commit/de9edffc25bc25741d1a856b52071c330c36f15b))


### Chores

* **cli:** let `--format raw` be used in conjunction with `--transform` ([1dfbfc5](https://github.com/trycourier/courier-cli/commit/1dfbfc55e22eebc1944bb3d11848f094311ee092))
* **internal:** regenerate SDK with no functional changes ([8c16167](https://github.com/trycourier/courier-cli/commit/8c161672c1fc1cf23c6cd83f5534b845807098e6))
* mark all CLI-related tests in Go with `t.Parallel()` ([3fb3d21](https://github.com/trycourier/courier-cli/commit/3fb3d215383204aa49743d2be374636da370afa6))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([cf2b66d](https://github.com/trycourier/courier-cli/commit/cf2b66de2f94d4dbc2dc391e116a6f71548939f8))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([e6cd02b](https://github.com/trycourier/courier-cli/commit/e6cd02bbbbbaf8b7ff84f948d9483977c06cd30a))

## 3.3.0 (2026-04-02)

Full Changelog: [v3.2.0...v3.3.0](https://github.com/trycourier/courier-cli/compare/v3.2.0...v3.3.0)

### Features

* add default description for enum CLI flags without an explicit description ([823fab6](https://github.com/trycourier/courier-cli/commit/823fab6fa75bcfc8ee2e8d5dca70b2f7db36a967))
* **api:** add create/retrieve/archive/publish/replace methods, event-id param to notifications ([f586c70](https://github.com/trycourier/courier-cli/commit/f586c70a88de4be8995c0ef384eefe56a326e1de))
* **api:** add providers CRUD and catalog list endpoints ([63678ee](https://github.com/trycourier/courier-cli/commit/63678ee1d1186dae6503dca94c243144d366b470))
* **api:** add routing-strategies resource ([b890b18](https://github.com/trycourier/courier-cli/commit/b890b182dece67bd2c4d7f85ec800e22bc59c1bd))
* set CLI flag constant values automatically where `x-stainless-const` is set ([614f808](https://github.com/trycourier/courier-cli/commit/614f808e7334e3adae68d5496319a666f50be09c))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([492dabf](https://github.com/trycourier/courier-cli/commit/492dabffb192428fcd123625cdbc5272c7aee16a))
* better support passing client args in any position ([1dad6ae](https://github.com/trycourier/courier-cli/commit/1dad6ae70c2fa1647dc1ba5fce44c53bef9fd00a))
* cli no longer hangs when stdin is attached to a pipe with empty input ([5fe52d1](https://github.com/trycourier/courier-cli/commit/5fe52d1a064f7c4b2346bbacc350ac737ead2081))
* fix for off-by-one error in pagination logic ([07c0b32](https://github.com/trycourier/courier-cli/commit/07c0b323427405e86bc80b01d489e79b1b6db3d5))
* handle empty data set using `--format explore` ([4f47249](https://github.com/trycourier/courier-cli/commit/4f4724971e9e79203aecada10d2f04316270ae70))
* improve linking behavior when developing on a branch not in the Go SDK ([4dd8baf](https://github.com/trycourier/courier-cli/commit/4dd8bafb13bed2cf62ff5f3ea77e92a151d11154))
* improved workflow for developing on branches ([f99f909](https://github.com/trycourier/courier-cli/commit/f99f909d00174b9325924e53d23c8e0939d2b1b7))
* no longer require an API key when building on production repos ([d125b34](https://github.com/trycourier/courier-cli/commit/d125b3498d30b59ee7c28a062ff54c3df942e9e0))
* only set client options when the corresponding CLI flag or env var is explicitly set ([a428dd5](https://github.com/trycourier/courier-cli/commit/a428dd5b300cb862478eab3a6f40cbe30af4a9bd))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([9894736](https://github.com/trycourier/courier-cli/commit/98947362639207cd34f9bb54f1841f895b06fbb0))


### Chores

* **ci:** skip lint on metadata-only changes ([b5909a3](https://github.com/trycourier/courier-cli/commit/b5909a34bd3d5322acb1d0d3da42c8907372e12a))
* **internal:** regenerate SDK with no functional changes ([c40831e](https://github.com/trycourier/courier-cli/commit/c40831e9e4a5fd779811dea448c1871ff337ecd1))
* **internal:** regenerate SDK with no functional changes ([310af9f](https://github.com/trycourier/courier-cli/commit/310af9fe454d1f9d69021cb9a505e6990899bdcb))
* **internal:** tweak CI branches ([8302597](https://github.com/trycourier/courier-cli/commit/83025978a62b26ceac906e634c3e68af4ee9f1ad))
* **internal:** update gitignore ([fd4e24e](https://github.com/trycourier/courier-cli/commit/fd4e24e16d3cd271b9b75121b38b792f86490e33))
* omit full usage information when missing required CLI parameters ([f4fe315](https://github.com/trycourier/courier-cli/commit/f4fe3151a945021ad870b883fe1c7fe00a19b4ad))


### Documentation

* add AGENTS.md for AI coding assistants ([b846ec0](https://github.com/trycourier/courier-cli/commit/b846ec0c161d736684ac085157bb56a2b896aaf0))
* document next branch for Stainless and CI parity ([94b0b79](https://github.com/trycourier/courier-cli/commit/94b0b79e0d2c4eeb5d918ffd8cf7e21e6b998d5a))

## 3.2.0 (2026-03-12)

Full Changelog: [v3.1.6...v3.2.0](https://github.com/trycourier/courier-cli/compare/v3.1.6...v3.2.0)

### Features

* [feat] add journeys api ([ad19639](https://github.com/trycourier/courier-cli/commit/ad196390441597a2bd3a754468dbd0cfa9dfc120))
* add `--max-items` flag for paginated/streaming endpoints ([1b04d3a](https://github.com/trycourier/courier-cli/commit/1b04d3ac7ba4296f61e0b0c857e2275032b0dbe0))
* support passing required body params through pipes ([a555982](https://github.com/trycourier/courier-cli/commit/a555982030cb1767e781eeeb88c72a4de084fde9))


### Bug Fixes

* fix for encoding arrays with `any` type items ([e5d46ce](https://github.com/trycourier/courier-cli/commit/e5d46ce9dfc8fce75ad022d7056b9844d0aed4c2))
* fix for test cases with newlines in YAML and better error reporting ([58be601](https://github.com/trycourier/courier-cli/commit/58be601bf0eb0189a091a3da87bbfbcf69efff4d))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([7d89ba3](https://github.com/trycourier/courier-cli/commit/7d89ba38e9ff375908207f077ae745123343fa11))
* **internal:** codegen related update ([05da3b8](https://github.com/trycourier/courier-cli/commit/05da3b8bce65dc6b4418aafa51e163bbc29e8275))
* **internal:** regenerate SDK with no functional changes ([ba5bcbe](https://github.com/trycourier/courier-cli/commit/ba5bcbe571ab9a11c640e7d1546efae702fe5d2d))
* **internal:** remove Homebrew distribution support ([53995d5](https://github.com/trycourier/courier-cli/commit/53995d50cd79d160eee41c9e1f858a104dbaa1d9))
* remove custom code ([4f5ec1b](https://github.com/trycourier/courier-cli/commit/4f5ec1bfba96d8d9c36125f7b99b9af1e100f15c))
* restore custom npm distribution and README patches ([404d643](https://github.com/trycourier/courier-cli/commit/404d6438fc8c38236624bbb1c94a0b4ab4dabaee))

## 3.1.6 (2026-03-04)

Full Changelog: [v3.1.5...v3.1.6](https://github.com/trycourier/courier-cli/compare/v3.1.5...v3.1.6)

### Bug Fixes

* add Node shim so npm creates bin symlink on install ([b1d475e](https://github.com/trycourier/courier-cli/commit/b1d475efd1672a48b5e6fdde1af4c516475b7bec))

## 3.1.5 (2026-03-04)

Full Changelog: [v3.1.4...v3.1.5](https://github.com/trycourier/courier-cli/compare/v3.1.4...v3.1.5)

### Documentation

* add README for npm package page ([823c5eb](https://github.com/trycourier/courier-cli/commit/823c5eb0591f55a0255f5943bebbe7b4c5075d85))

## 3.1.4 (2026-03-03)

Full Changelog: [v3.1.3...v3.1.4](https://github.com/trycourier/courier-cli/compare/v3.1.3...v3.1.4)

### Bug Fixes

* update npm for OIDC trusted publishing support ([71f7928](https://github.com/trycourier/courier-cli/commit/71f792878fb739fcc9d328ac217ac88fe3b1d04d))

## 3.1.3 (2026-03-03)

Full Changelog: [v3.1.2...v3.1.3](https://github.com/trycourier/courier-cli/compare/v3.1.2...v3.1.3)

### Bug Fixes

* **api:** remove duplicate token parameter from users tokens add-single ([be432e1](https://github.com/trycourier/courier-cli/commit/be432e12162da46bdab567473b988bddcb39adca))
* switch npm publish to OIDC trusted publishing ([a01b06f](https://github.com/trycourier/courier-cli/commit/a01b06fca86221c7c9a883ee98ff459913b2d2f1))

## 3.1.2 (2026-03-03)

Full Changelog: [v3.1.1...v3.1.2](https://github.com/trycourier/courier-cli/compare/v3.1.1...v3.1.2)

### Bug Fixes

* use correct NPM_AUTH_TOKEN secret name for npm publish ([a291c22](https://github.com/trycourier/courier-cli/commit/a291c22c1f7e5d6d4596c6cfa0c443e52f020863))

## 3.1.1 (2026-03-03)

Full Changelog: [v3.1.0...v3.1.1](https://github.com/trycourier/courier-cli/compare/v3.1.0...v3.1.1)

### Bug Fixes

* decouple npm publish from goreleaser success ([b8b4e71](https://github.com/trycourier/courier-cli/commit/b8b4e7156836daa210bed3116bea170f119a091b))

## 3.1.0 (2026-03-03)

Full Changelog: [v3.0.0...v3.1.0](https://github.com/trycourier/courier-cli/compare/v3.0.0...v3.1.0)

### Features

* add npm distribution wrapper for Go CLI binary ([5756f2c](https://github.com/trycourier/courier-cli/commit/5756f2ce7c9d2373ade8545ab0ce94bab40a74d4))


### Chores

* **internal:** add Homebrew distribution support ([5fee3a3](https://github.com/trycourier/courier-cli/commit/5fee3a3383e57665b7be7daca582980629c749c6))

## 3.0.0 (2026-03-02)

Full Changelog: [v3.0.0...v3.0.0](https://github.com/trycourier/courier-cli/compare/v3.0.0...v3.0.0)

### Bug Fixes

* **cli:** set go_sdk_package to resolve v4 module path ([c8b39bd](https://github.com/trycourier/courier-cli/commit/c8b39bd082497d823d6872cbc967ee7426610c33))


### Chores

* sync repo ([e6a9cbd](https://github.com/trycourier/courier-cli/commit/e6a9cbd2015d7c7b6d96688b0f81210bc4e22f63))

## 3.0.0 (2026-03-02)

Full Changelog: [v0.0.1...v3.0.0](https://github.com/trycourier/courier-cli/compare/v0.0.1...v3.0.0)

### Bug Fixes

* **cli:** set go_sdk_package to resolve v4 module path ([c8b39bd](https://github.com/trycourier/courier-cli/commit/c8b39bd082497d823d6872cbc967ee7426610c33))


### Chores

* sync repo ([e6a9cbd](https://github.com/trycourier/courier-cli/commit/e6a9cbd2015d7c7b6d96688b0f81210bc4e22f63))
