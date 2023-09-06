package ldconsul

import (
	"testing"

	c "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"

	"github.com/launchdarkly/go-sdk-common/v3/ldlog"
	"github.com/launchdarkly/go-sdk-common/v3/ldlogtest"
	"github.com/launchdarkly/go-server-sdk/v7/subsystems"
	"github.com/launchdarkly/go-server-sdk/v7/testhelpers/storetest"
)

func TestConsulDataStore(t *testing.T) {
	storetest.NewPersistentDataStoreTestSuite(makeTestStore, clearTestData).
		ErrorStoreFactory(makeFailedStore(), verifyFailedStoreError).
		ConcurrentModificationHook(setConcurrentModificationHook).
		Run(t)
}

func TestLoggingAtStartup(t *testing.T) {
	expectAddress := func(t *testing.T, builder *DataStoreBuilder, message string) {
		mockLog := ldlogtest.NewMockLog()
		ctx := subsystems.BasicClientContext{}
		ctx.Logging.Loggers = mockLog.Loggers
		store, _ := builder.Build(ctx)
		defer store.Close()
		mockLog.AssertMessageMatch(t, true, ldlog.Info, message)
	}

	t.Run("default address", func(t *testing.T) {
		expectAddress(t, DataStore(),
			"ConsulDataStore: Using Consul server at 127.0.0.1:8500")
	})

	t.Run("custom address", func(t *testing.T) {
		expectAddress(t, DataStore().Address("myhost:1000"),
			"ConsulDataStore: Using Consul server at myhost:1000")
	})
}

func makeTestStore(prefix string) subsystems.ComponentConfigurer[subsystems.PersistentDataStore] {
	return DataStore().Prefix(prefix)
}

func makeFailedStore() subsystems.ComponentConfigurer[subsystems.PersistentDataStore] {
	// Here we ensure that all Consul operations will fail by providing an invalid hostname.
	return DataStore().Address("not-a-real-consul-host")
}

func verifyFailedStoreError(t assert.TestingT, err error) {
	assert.Contains(t, err.Error(), "no such host")
}

func clearTestData(prefix string) error {
	client, err := c.NewClient(c.DefaultConfig())
	if err != nil {
		return err
	}
	kv := client.KV()
	_, err = kv.DeleteTree(prefix+"/", nil)
	return err
}

func setConcurrentModificationHook(store subsystems.PersistentDataStore, hook func()) {
	store.(*consulDataStoreImpl).testTxHook = hook
}
