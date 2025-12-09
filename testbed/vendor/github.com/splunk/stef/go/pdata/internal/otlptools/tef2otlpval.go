package otlptools

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/splunk/stef/go/otel/oteltef"
)

//func TefToOtlpMap(in *oteltef.Attributes, out pcommon.Map) error {
//	out.EnsureCapacity(in.Len())
//
//	decoder := anyvalue.Decoder{}
//	for i := 0; i < in.Len(); i++ {
//		kv := in.At(i)
//		val := out.PutEmpty(kv.Key())
//		decoder.Reset(anyvalue.ImmutableBytes(kv.Value().Bytes()))
//		err := tefToOtlpValue(&decoder, val)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func TefToOtlpMap(in *oteltef.Attributes, out pcommon.Map) error {
	out.EnsureCapacity(in.Len())

	//decoder := anyvalue.Decoder{}
	for i := 0; i < in.Len(); i++ {
		kv := in.At(i)
		val := out.PutEmpty(kv.Key())
		//decoder.Reset(anyvalue.ImmutableBytes(kv.Value()))
		err := tefAnyValueToOtlp(kv.Value(), val)
		if err != nil {
			return err
		}
	}
	return nil
}

var errDecode = errors.New("decode error")

func tefAnyValueToOtlp(anyVal *oteltef.AnyValue, into pcommon.Value) error {
	switch anyVal.Type() {
	case oteltef.AnyValueTypeString:
		into.SetStr(anyVal.String())

	case oteltef.AnyValueTypeBytes:
		bytes := into.SetEmptyBytes()
		bytes.Append([]byte(anyVal.Bytes())...)

	case oteltef.AnyValueTypeInt64:
		into.SetInt(anyVal.Int64())

	case oteltef.AnyValueTypeBool:
		into.SetBool(anyVal.Bool())

	case oteltef.AnyValueTypeNone:

	case oteltef.AnyValueTypeFloat64:
		into.SetDouble(anyVal.Float64())

	case oteltef.AnyValueTypeArray:
		values := into.SetEmptySlice()
		arr := anyVal.Array()
		for i := 0; i < arr.Len(); i++ {
			val := values.AppendEmpty()
			err := tefAnyValueToOtlp(arr.At(i), val)
			if err != nil {
				return err
			}
		}

	case oteltef.AnyValueTypeKVList:
		values := into.SetEmptyMap()
		kvList := anyVal.KVList()
		for i := 0; i < kvList.Len(); i++ {
			pair := kvList.At(i)
			val := values.PutEmpty(pair.Key())
			err := tefAnyValueToOtlp(pair.Value(), val)
			if err != nil {
				return err
			}
		}

	default:
		return errDecode
	}
	return nil
}
