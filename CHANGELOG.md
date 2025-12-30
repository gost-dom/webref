# Changelog - Gost-DOM Webref


## [0.5.13](https://github.com/gost-dom/webref/compare/v0.5.12...v0.5.13) (2025-12-30)

## [0.5.12](https://github.com/gost-dom/webref/compare/v0.5.11...v0.5.12) (2025-12-29)


### Features

* Implement Inheritance for Dictionary types ([3ba19ba](https://github.com/gost-dom/webref/commit/3ba19baa33618efbe5c001e160fbbaed79a3f475))

## [0.5.11](https://github.com/gost-dom/webref/compare/v0.5.10...v0.5.11) (2025-12-17)


### Features

* Export Exposed extended attribute on interfaces ([9435c98](https://github.com/gost-dom/webref/commit/9435c98fda20204c3091f5d039a7e3078e66ca2b))
* Export interface Global extended attributes ([1d94679](https://github.com/gost-dom/webref/commit/1d946794fd5b4db78bcf2ae70f78497b26b40396))

## [0.5.10](https://github.com/gost-dom/webref/compare/v0.5.9...v0.5.10) (2025-11-18)


### Features

* Support partial interfaces, and merging ([8242dcc](https://github.com/gost-dom/webref/commit/8242dcc7aca8f0f7e343710f4066d1eefe1b4199))

## [0.5.9](https://github.com/gost-dom/webref/compare/v0.5.8...v0.5.9) (2025-07-03)


### Features

* Expose dictionary specifications ([f100a8a](https://github.com/gost-dom/webref/commit/f100a8a120f57027bca7190abfc77cced5afdb3e))

## [0.5.8](https://github.com/gost-dom/webref/compare/v0.5.7...v0.5.8) (2025-06-29)


### Features

* Support Promise<> return types ([666a361](https://github.com/gost-dom/webref/commit/666a3619df580b7737ae023894e92f3e9f98cd27))

## [0.5.7](https://github.com/gost-dom/webref/compare/v0.5.6...v0.5.7) (2025-06-16)


### Features

* Expose argument default values ([00c35c9](https://github.com/gost-dom/webref/commit/00c35c9bef6633107e2a06bb9a478b270f0a305e))

## [0.5.6](https://github.com/gost-dom/webref/compare/v0.5.5...v0.5.6) (2025-05-24)


### Bug Fixes

* Defer file close _after_ checking for error ([db5ec46](https://github.com/gost-dom/webref/commit/db5ec4615c1994cc3ea7b5f86959afe2c81d1cdc))

## [0.5.5](https://github.com/gost-dom/webref/compare/v0.5.4...v0.5.5) (2025-05-21)


### Bug Fixes

* union types can be nullable ([5a3db2c](https://github.com/gost-dom/webref/commit/5a3db2c938fca276602c7a6734f00d4ff6595870))

## [0.5.4](https://github.com/gost-dom/webref/compare/v0.5.3...v0.5.4) (2025-05-21)


### Features

* Support union types in operation arguments ([2e5a6c2](https://github.com/gost-dom/webref/commit/2e5a6c2abb38bc77290d9835af7496e4d1d1e14e))

## [0.5.3](https://github.com/gost-dom/webref/compare/v0.5.2...v0.5.3) (2025-05-17)


### Features

* Expose iterable types in interface ([d473db1](https://github.com/gost-dom/webref/commit/d473db16dc5f5d95b0c83e6cce14b52b2566c669))

## [0.5.2](https://github.com/gost-dom/webref/compare/v0.5.1...v0.5.2) (2025-04-09)


### Features

* Add Interface.Constructors property ([ee63d82](https://github.com/gost-dom/webref/commit/ee63d82a37789160783e2f8b8b5f9eed975db694))

## [0.5.1](https://github.com/gost-dom/webref/compare/v0.5.0...v0.5.1) (2025-03-02)


### Bug Fixes

* Remove Print statement ([c7a66b2](https://github.com/gost-dom/webref/commit/c7a66b2289ec6b19e70188a5b907e30ebc69386c))

## [0.5.0](https://github.com/gost-dom/webref/compare/v0.4.0...v0.5.0) (2025-03-02)


### ⚠ BREAKING CHANGES

* Move the legacy JSON decoded types to a `legacy`
subpackage for easy identifaction of APIs to avoid unless it's the only
source of that information. See readme file for more information.

* Move legacy idl data to `legacy` sub package. ([649e529](https://github.com/gost-dom/webref/commit/649e529e4bf965e2d604979cebae526158c9a5ba))

## [0.4.0](https://github.com/gost-dom/webref/compare/v0.3.1...v0.4.0) (2025-03-02)


### ⚠ BREAKING CHANGES

* This is a breaking change for v0.3.1, but not v0.3.0 -
as 0.3.1 introduced the new type but with poor names that is improved in
v0.4.0.

### Features

* Fix poorly named exported fields and types ([d20d877](https://github.com/gost-dom/webref/commit/d20d877adfba6089a9689e0f37dcb01d5f56f54a))

## [0.3.1](https://github.com/gost-dom/webref/compare/v0.3.0...v0.3.1) (2025-03-02)

> [!WARNING]
>
> Avoid this version, it has poorly named fields, that are improved in v0.4.0

### Features

* Implement reading event data ([76db27c](https://github.com/gost-dom/webref/commit/76db27c555f85c00dd87d10665d40af171b9b813))

## [0.2.1](https://github.com/gost-dom/webref/compare/v0.2.0...v0.2.1) (2025-03-02)


### Features

* Implement reading event data ([76db27c](https://github.com/gost-dom/webref/commit/76db27c555f85c00dd87d10665d40af171b9b813))
* Support optional arguments ([a4a9b67](https://github.com/gost-dom/webref/commit/a4a9b6718e0d23ed2e4e855c7d69944b2365f63d))

## [0.3.0](https://github.com/gost-dom/webref/compare/v0.2.0...v0.2.1) (2025-03-02)

### Features

* Support optional arguments ([a4a9b67](https://github.com/gost-dom/webref/commit/a4a9b6718e0d23ed2e4e855c7d69944b2365f63d))

## [0.2.0](https://github.com/gost-dom/webref/compare/v0.1.0...v0.2.0) (2025-02-19)

### Features

* Handle sequence types ([10c91ab](https://github.com/gost-dom/webref/commit/10c91ab7df467d4cc9ee42354012c73fb68ee681))
