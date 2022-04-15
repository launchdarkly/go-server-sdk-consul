# Change log

All notable changes to the LaunchDarkly Go SDK Consul integration will be documented in this file. This project adheres to [Semantic Versioning](http://semver.org).

## [1.0.2] - 2022-04-15
### Fixed:
- Updated to v1.12.0 of the Consul API client to prevent a vulnerability warning. ([#9](https://github.com/launchdarkly/go-server-sdk-consul/issues/9))

## [1.0.1] - 2021-11-16
### Changed:
- Updated the dependency version of `github.com/hashicorp/consul/api` to v1.11.0. This was to address vulnerabilities that have been reported against earlier versions of Consul. We believe that those CVE reports are somewhat misleading since they refer to the Consul _server_, rather than the API library, but vulnerability scanners often conflate the two and the only known workaround is to update the API version (see https://github.com/hashicorp/consul/issues/10674).

## [1.0.0] - 2020-09-18
Initial release of the stand-alone version of this package to be used with versions 5.0.0 and above of the LaunchDarkly Server-Side SDK for Go.
