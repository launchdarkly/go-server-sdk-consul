package ldconsul

import (
	"testing"

	c "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"

	"gopkg.in/launchdarkly/go-sdk-common.v2/ldlog"
	"gopkg.in/launchdarkly/go-sdk-common.v2/ldlogtest"
	"gopkg.in/launchdarkly/go-server-sdk.v5/interfaces"
	"gopkg.in/launchdarkly/go-server-sdk.v5/ldcomponents"
	"gopkg.in/launchdarkly/go-server-sdk.v5/testhelpers"
	"gopkg.in/launchdarkly/go-server-sdk.v5/testhelpers/storetest"
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
		ctx := testhelpers.NewSimpleClientContext("").WithLogging(
			ldcomponents.Logging().Loggers(mockLog.Loggers),
		)
		store, _ := builder.CreatePersistentDataStore(ctx)
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

func makeTestStore(prefix string) interfaces.PersistentDataStoreFactory {
	return DataStore().Prefix(prefix)
}

func makeFailedStore() interfaces.PersistentDataStoreFactory {
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

func setConcurrentModificationHook(store interfaces.PersistentDataStore, hook func()) {
	store.(*consulDataStoreImpl).testTxHook = hook
}
