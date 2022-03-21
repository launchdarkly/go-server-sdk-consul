package ldconsul

import (
	"testing"

	c "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/launchdarkly/go-sdk-common/v3/ldvalue"
	"github.com/launchdarkly/go-server-sdk/v6/interfaces"
	"github.com/launchdarkly/go-server-sdk/v6/ldcomponents"
)

func TestDataStoreBuilder(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		b := DataStore()
		assert.Equal(t, c.Config{}, b.consulConfig)
		assert.Equal(t, DefaultPrefix, b.prefix)
	})

	t.Run("Address", func(t *testing.T) {
		b := DataStore().Address("a")
		assert.Equal(t, "a", b.consulConfig.Address)
	})

	t.Run("Config", func(t *testing.T) {
		var config c.Config
		config.Address = "a"

		b := DataStore().Config(config)
		assert.Equal(t, config, b.consulConfig)
	})

	t.Run("Prefix", func(t *testing.T) {
		b := DataStore().Prefix("p")
		assert.Equal(t, "p", b.prefix)

		b.Prefix("")
		assert.Equal(t, DefaultPrefix, b.prefix)
	})

	t.Run("error for invalid address", func(t *testing.T) {
		b := DataStore().Address("bad-scheme://no")
		_, err := b.CreatePersistentDataStore(simpleTestContext{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Unknown protocol")
	})

	t.Run("describe configuration", func(t *testing.T) {
		assert.Equal(t, ldvalue.String("Consul"), DataStore().DescribeConfiguration())
	})
}

// stub implementation of interfaces.ClientContext
type simpleTestContext struct{}

func (c simpleTestContext) GetBasic() interfaces.BasicConfiguration {
	return interfaces.BasicConfiguration{}
}

func (c simpleTestContext) GetHTTP() interfaces.HTTPConfiguration {
	return nil
}

func (c simpleTestContext) GetLogging() interfaces.LoggingConfiguration {
	lc, _ := ldcomponents.Logging().CreateLoggingConfiguration(interfaces.BasicConfiguration{})
	return lc
}
