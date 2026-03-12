# Changelog

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
