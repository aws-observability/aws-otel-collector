package awscontainerinsightreceiver

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(t, err)

	factory := NewFactory()
	receiverType := "awscontainerinsightreceiver"
	factories.Receivers[config.Type(receiverType)] = factory
	cfg, err := configtest.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, len(cfg.Receivers), 2)

	//ensure default configurations are generated when users provide nothing
	r0 := cfg.Receivers[receiverType]
	assert.Equal(t, factory.CreateDefaultConfig(), r0)

	r1 := cfg.Receivers["awscontainerinsightreceiver"]
	assert.Equal(t, r1, factory.CreateDefaultConfig())

	r2 := cfg.Receivers["awscontainerinsightreceiver/collection_interval_settings"].(*Config)
	assert.Equal(t, r2,
		&Config{
			ReceiverSettings: config.ReceiverSettings{
				TypeVal: config.Type(receiverType),
				NameVal: "awscontainerinsightreceiver/collection_interval_settings",
			},
			CollectionInterval:    60 * time.Second,
			ContainerOrchestrator: "eks",
		})
}
