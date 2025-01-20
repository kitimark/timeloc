# Changelog

## [1.1.0](https://github.com/kitimark/timeloc/compare/v1.0.0...v1.1.0) (2025-01-18)


### Features

* **gclplugin:** add timeloc plugin for golangci-lint integration ([#6](https://github.com/kitimark/timeloc/issues/6)) ([82b5802](https://github.com/kitimark/timeloc/commit/82b58028c4b255d954ff78e4047681fd0f2090d5))


### Bug Fixes

* **timeloc:** detect time method chaining and simplify error message ([#9](https://github.com/kitimark/timeloc/issues/9)) ([ebc9e48](https://github.com/kitimark/timeloc/commit/ebc9e48462942894cbf0c45883f300937431de7a))

## 1.0.0 (2025-01-17)


### Features

* initiate timeloc ([9f74683](https://github.com/kitimark/timeloc/commit/9f74683fc15fea4140cd796304355a4e28840b4d))
* **timeloc:** not allowed time.Local usage ([#1](https://github.com/kitimark/timeloc/issues/1)) ([9dc1b44](https://github.com/kitimark/timeloc/commit/9dc1b447f617def2329a58e2faefec965133ed35))


### Bug Fixes

* **timeloc:** detect false positive when time package is aliased ([#5](https://github.com/kitimark/timeloc/issues/5)) ([ad2e216](https://github.com/kitimark/timeloc/commit/ad2e2168388186d660cdd32592297376dc0ad3f0))
* **timeloc:** program will panic, when function have a built-in interface in argument ([be6b6c3](https://github.com/kitimark/timeloc/commit/be6b6c314ae399d9d833b75db91835d7eca4c793))
