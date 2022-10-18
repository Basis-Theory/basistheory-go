# [3.7.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.6.0...v3.7.0) (2022-10-18)


### Features

* adds support for access rules and token containers ([#53](https://github.com/Basis-Theory/basistheory-go/issues/53)) ([cc73b75](https://github.com/Basis-Theory/basistheory-go/commit/cc73b7507a91cf56bafd650d48737dffde4f2134))

# [3.6.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.5.0...v3.6.0) (2022-08-10)


### Features

* update app types and add tokens container field ([#51](https://github.com/Basis-Theory/basistheory-go/issues/51)) ([8869091](https://github.com/Basis-Theory/basistheory-go/commit/8869091236381d7b8905d482ee726ef04755a607))

# [3.5.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.4.0...v3.5.0) (2022-08-01)


### Features

* add token expires_at, validations, updated models  ([#45](https://github.com/Basis-Theory/basistheory-go/issues/45)) ([28b2b79](https://github.com/Basis-Theory/basistheory-go/commit/28b2b79a5a1774c500619409630fe8a7f11c78d8))

# [3.4.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.3.1...v3.4.0) (2022-07-21)


### Features

* removing atomics ([#44](https://github.com/Basis-Theory/basistheory-go/issues/44)) ([8dd7e11](https://github.com/Basis-Theory/basistheory-go/commit/8dd7e11bc64c2e84f1f1f0ca8362fc64515798eb))

## [3.3.1](https://github.com/Basis-Theory/basistheory-go/compare/v3.3.0...v3.3.1) (2022-06-30)


### Code Refactoring

* remove atomic react methods, fix required params ([#43](https://github.com/Basis-Theory/basistheory-go/issues/43)) ([a4eef9c](https://github.com/Basis-Theory/basistheory-go/commit/a4eef9c95cd11862b333f4d277a635f73367d0e5))


### BREAKING CHANGES

* remove react methods from atomic banks and cards clients

# [3.3.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.2.0...v3.3.0) (2022-06-30)


### Features

* add id to create token request and return id string on responses ([#42](https://github.com/Basis-Theory/basistheory-go/issues/42)) ([7d3d581](https://github.com/Basis-Theory/basistheory-go/commit/7d3d5818f7043b3dd5928d0a28cb515e51a2efcc))

# [3.2.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.1.0...v3.2.0) (2022-06-18)


### Features

* adds PATCH on tokens and mask property ([#41](https://github.com/Basis-Theory/basistheory-go/issues/41)) ([03a2abe](https://github.com/Basis-Theory/basistheory-go/commit/03a2abea80f76a037d867f346daf076a12e162e8))

# [3.1.0](https://github.com/Basis-Theory/basistheory-go/compare/v3.0.0...v3.1.0) (2022-06-14)


### Features

* adds deduplicate_token on Tokens. adds settings on Tenants ([#40](https://github.com/Basis-Theory/basistheory-go/issues/40)) ([ec1b574](https://github.com/Basis-Theory/basistheory-go/commit/ec1b574b146bcaf331d97e0c6329b7b04d9179c6))

# [3.0.0](https://github.com/Basis-Theory/basistheory-go/compare/v2.0.1...v3.0.0) (2022-06-10)


* feat!: adds Tokens, Tokenize, Logs, Permissions, and other endpoints (#38) ([f79df19](https://github.com/Basis-Theory/basistheory-go/commit/f79df19fc0b68f8c3d574dae356c7e3835b55383)), closes [#38](https://github.com/Basis-Theory/basistheory-go/issues/38)


### BREAKING CHANGES

* updates CRUD function names for cleaner interface

## [2.0.1](https://github.com/Basis-Theory/basistheory-go/compare/v2.0.0...v2.0.1) (2022-06-03)


### Bug Fixes

* fixes go.mod module for v2 ([#36](https://github.com/Basis-Theory/basistheory-go/issues/36)) ([b0d5721](https://github.com/Basis-Theory/basistheory-go/commit/b0d5721af398d189ffd026b4f1c25a426869b45a))

# [2.0.0](https://github.com/Basis-Theory/basistheory-go/compare/v1.1.0...v2.0.0) (2022-06-03)


* feat!: adds Proxy resource, and Application on Reactor (#35) ([5d0f021](https://github.com/Basis-Theory/basistheory-go/commit/5d0f0215354a9218f10c993e524547dc7d0ca30b)), closes [#35](https://github.com/Basis-Theory/basistheory-go/issues/35)


### BREAKING CHANGES

* updates API, function, and model names to match new swagger spec

# [1.1.0](https://github.com/Basis-Theory/basistheory-go/compare/v1.0.0...v1.1.0) (2022-04-15)


### Features

* removes return model from reactor-formula delete ([#33](https://github.com/Basis-Theory/basistheory-go/issues/33)) ([fca0c77](https://github.com/Basis-Theory/basistheory-go/commit/fca0c7772809b6978226ca0870c5daa5738845b7))

# 1.0.0 (2022-04-14)


### Features

* initial commit ([50b09a1](https://github.com/Basis-Theory/basistheory-go/commit/50b09a1256d9f96a71984d37e88f6f065edb115b))
