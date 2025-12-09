package otlptools

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type AttrAccessible struct {
	Key   string
	Value pcommon.Value
}

func Map2attrs(attributes pcommon.Map) []AttrAccessible {
	attrs := make([]AttrAccessible, 0, attributes.Len())
	attributes.Range(
		func(k string, v pcommon.Value) bool {
			attrs = append(attrs, AttrAccessible{Key: k, Value: v})
			return true
		},
	)
	return attrs
}

func CmpResourceMetrics(left, right pmetric.ResourceMetrics) int {
	c := strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}
	return CmpAttrs(left.Resource().Attributes(), right.Resource().Attributes())
}

func CmpResourceSpans(left, right ptrace.ResourceSpans) int {
	c := strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}
	return CmpAttrs(left.Resource().Attributes(), right.Resource().Attributes())
}

func CmpScopeMetrics(left, right pmetric.ScopeMetrics) int {
	c := strings.Compare(left.Scope().Name(), right.Scope().Name())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.Scope().Version(), right.Scope().Version())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}

	return CmpAttrs(left.Scope().Attributes(), right.Scope().Attributes())
}

func CmpScopeSpans(left, right ptrace.ScopeSpans) int {
	c := strings.Compare(left.Scope().Name(), right.Scope().Name())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.Scope().Version(), right.Scope().Version())
	if c != 0 {
		return c
	}
	c = strings.Compare(left.SchemaUrl(), right.SchemaUrl())
	if c != 0 {
		return c
	}

	return CmpAttrs(left.Scope().Attributes(), right.Scope().Attributes())
}

func CmpAttrs(a, b pcommon.Map) int {
	left := Map2attrs(a)
	right := Map2attrs(b)

	lenDiff := len(left) - len(right)
	l := min(len(left), len(right))
	for i := 0; i < l; i++ {
		c := strings.Compare(left[i].Key, right[i].Key)
		if c != 0 {
			return c
		}
	}
	if lenDiff != 0 {
		return lenDiff
	}
	for i := 0; i < l; i++ {
		c := CmpVal(left[i].Value, right[i].Value)
		if c != 0 {
			return c
		}
	}
	return lenDiff
}

func CmpVal(left, right pcommon.Value) int {
	c := int(left.Type()) - int(right.Type())
	if c != 0 {
		return c
	}

	switch left.Type() {
	case pcommon.ValueTypeStr:
		return strings.Compare(left.Str(), right.Str())

	case pcommon.ValueTypeInt:
		return CmpInt64(left.Int(), right.Int())

	case pcommon.ValueTypeBool:
		return CmpBool(left.Bool(), right.Bool())

	case pcommon.ValueTypeSlice:
		left := left.Slice()
		right := right.Slice()
		if left.Len() != right.Len() {
			return left.Len() - right.Len()
		}
		for i := 0; i < left.Len(); i++ {
			c := CmpVal(left.At(i), right.At(i))
			if c != 0 {
				return c
			}
		}
		return 0

	case pcommon.ValueTypeEmpty:
		return 0

	default:
		panic("comparison not implemented")
	}
}

func CmpBool(v1, v2 bool) int {
	if v1 == v2 {
		return 0
	}
	if v1 {
		return 1
	}
	return -1
}

func CmpInt64(v1, v2 int64) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
}
