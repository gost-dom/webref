# Changelog - Gost-DOM Webref


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
