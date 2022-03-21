module github.com/launchdarkly/go-server-sdk-consul/v2

go 1.16

require (
	github.com/hashicorp/consul/api v1.11.0
	github.com/launchdarkly/go-sdk-common/v3 v3.0.0
	github.com/launchdarkly/go-server-sdk/v6 v6.0.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/launchdarkly/go-sdk-common/v3 => github.com/launchdarkly/go-sdk-common-private/v3 v3.0.0-alpha.4

replace github.com/launchdarkly/go-sdk-events/v2 => github.com/launchdarkly/go-sdk-events-private/v2 v2.0.0-alpha.4

replace github.com/launchdarkly/go-server-sdk-evaluation/v2 => github.com/launchdarkly/go-server-sdk-evaluation-private/v2 v2.0.0-alpha.6

replace github.com/launchdarkly/go-server-sdk/v6 => github.com/launchdarkly/go-server-sdk-private/v6 v6.0.0-alpha.1
