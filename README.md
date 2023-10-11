# LaunchDarkly Server-side SDK for Go - Consul integration

[![Circle CI](https://circleci.com/gh/launchdarkly/go-server-sdk-consul.svg?style=shield)](https://circleci.com/gh/launchdarkly/go-server-sdk-consul) [![Documentation](https://img.shields.io/static/v1?label=go.dev&message=reference&color=00add8)](https://pkg.go.dev/github.com/launchdarkly/go-server-sdk-consul/v3)

This library provides a [Consul](https://www.consul.io/)-backed persistence mechanism (data store) for the [LaunchDarkly Go SDK](https://github.com/launchdarkly/go-server-sdk), replacing the default in-memory data store. It uses the standard [Consul Go client](https://github.com/hashicorp/consul).

This version of the library requires at least version 6.0.0 of the LaunchDarkly Go SDK; for versions of the library to use with earlier SDK versions, see the changelog.

The minimum Go version is 1.18.

For more information, see also: [Using Consul as a persistent feature store](https://docs.launchdarkly.com/sdk/features/storing-data/consul#go).

## Quick setup

This assumes that you have already installed the LaunchDarkly Go SDK.

1. Import the LaunchDarkly SDK packages and the package for this library:

```go
import (
    ld "github.com/launchdarkly/go-server-sdk/v7"
    "github.com/launchdarkly/go-server-sdk/v7/ldcomponents"
    ldconsul "github.com/launchdarkly/go-server-sdk-consul/v3"
)
```

2. When configuring your SDK client, add the Consul data store as a `PersistentDataStore`. You may specify any custom Consul options using the methods of `ConsulDataStoreBuilder`. For instance, to customize the Consul hostname:

```go
    var config ld.Config{}
    config.DataStore = ldcomponents.PersistentDataStore(
        ldconsul.DataStore().Address("my-consul-host"),
    )
```

By default, the store will try to connect to a local Consul instance on port 8500.

## Caching behavior

The LaunchDarkly SDK has a standard caching mechanism for any persistent data store, to reduce database traffic. This is configured through the SDK's `PersistentDataStoreBuilder` class as described the SDK documentation. For instance, to specify a cache TTL of 5 minutes:

```go
    var config ld.Config{}
    config.DataStore = ldcomponents.PersistentDataStore(
        ldconsul.DataStore(),
    ).CacheMinutes(5)
```

## Data size limitation

Consul does not support storing values greater than 512KB. Therefore, this integration will not work if the JSON representation of any feature flag or user segment exceeds that size.

To see the JSON representations of all flags and segments, query `https://app.launchdarkly.com/sdk/latest-all` with your SDK key in an `Authorization` header.

## LaunchDarkly overview

[LaunchDarkly](https://www.launchdarkly.com) is a feature management platform that serves trillions of feature flags daily to help teams build better software, faster. [Get started](https://docs.launchdarkly.com/docs/getting-started) using LaunchDarkly today!

## About LaunchDarkly

* LaunchDarkly is a continuous delivery platform that provides feature flags as a service and allows developers to iterate quickly and safely. We allow you to easily flag your features and manage them from the LaunchDarkly dashboard.  With LaunchDarkly, you can:
    * Roll out a new feature to a subset of your users (like a group of users who opt-in to a beta tester group), gathering feedback and bug reports from real-world use cases.
    * Gradually roll out a feature to an increasing percentage of users, and track the effect that the feature has on key metrics (for instance, how likely is a user to complete a purchase if they have feature A versus feature B?).
    * Turn off a feature that you realize is causing performance problems in production, without needing to re-deploy, or even restart the application with a changed configuration file.
    * Grant access to certain features based on user attributes, like payment plan (eg: users on the ‘gold’ plan get access to more features than users in the ‘silver’ plan). Disable parts of your application to facilitate maintenance, without taking everything offline.
* LaunchDarkly provides feature flag SDKs for a wide variety of languages and technologies. Read [our documentation](https://docs.launchdarkly.com/sdk) for a complete list.
* Explore LaunchDarkly
    * [launchdarkly.com](https://www.launchdarkly.com/ "LaunchDarkly Main Website") for more information
    * [docs.launchdarkly.com](https://docs.launchdarkly.com/  "LaunchDarkly Documentation") for our documentation and SDK reference guides
    * [apidocs.launchdarkly.com](https://apidocs.launchdarkly.com/  "LaunchDarkly API Documentation") for our API documentation
    * [blog.launchdarkly.com](https://blog.launchdarkly.com/  "LaunchDarkly Blog Documentation") for the latest product updates
