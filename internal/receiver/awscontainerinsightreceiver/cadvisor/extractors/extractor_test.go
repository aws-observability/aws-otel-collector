package extractors

import (
	"testing"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/stretchr/testify/assert"
)

func TestCAdvisorMetric_Merge(t *testing.T) {
	src := &CAdvisorMetric{fields: map[string]interface{}{"value1": 1, "value2": 2}, tags: map[string]string{common.Timestamp: "1586331559882"}}
	dest := &CAdvisorMetric{fields: map[string]interface{}{"value1": 3, "value3": 3}, tags: map[string]string{common.Timestamp: "1586331559973"}}
	src.Merge(dest)
	assert.Equal(t, 3, len(src.fields))
	assert.Equal(t, 1, src.fields["value1"].(int))
}
