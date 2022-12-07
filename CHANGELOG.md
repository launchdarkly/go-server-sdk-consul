# Change log

All notable changes to the LaunchDarkly Go SDK Consul integration will be documented in this file. This project adheres to [Semantic Versioning](http://semver.org).

## [2.0.1] - 2022-12-07
### Fixed:
- Updated SDK dependency to use v6.0.0 release.

## [2.0.0] - 2022-12-07
This release corresponds to the 6.0.0 release of the LaunchDarkly Go SDK. Any application code that is being updated to use the 6.0.0 SDK, and was using a 1.x version of `go-server-sdk-consul`, should now use a 2.x version instead.

There are no functional differences in the behavior of the Consul integration; the differences are only related to changes in the usage of interface types for configuration in the SDK.

## [1.0.2] - 2022-04-15
### Fixed:
- Updated to v1.12.0 of the Consul API client to prevent a vulnerability warning. ([#9](https://github.com/launchdarkly/go-server-sdk-consul/issues/9))

## [1.0.1] - 2021-11-16
### Changed:
- Updated the dependency version of `github.com/hashicorp/consul/api` to v1.11.0. This was to address vulnerabilities that have been reported against earlier versions of Consul. We believe that those CVE reports are somewhat misleading since they refer to the Consul _server_, rather than the API library, but vulnerability scanners often conflate the two and the only known workaround is to update the API version (see https://github.com/hashicorp/consul/issues/10674).

## [1.0.0] - 2020-09-18
Initial release of the stand-alone version of this package to be used with versions 5.0.0 and above of the LaunchDarkly Server-Side SDK for Go.
