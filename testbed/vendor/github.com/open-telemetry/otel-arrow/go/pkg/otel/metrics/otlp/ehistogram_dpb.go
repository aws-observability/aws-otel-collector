/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otlp

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"go.opentelemetry.io/collector/pdata/pmetric"

	arrowutils "github.com/open-telemetry/otel-arrow/go/pkg/arrow"
	"github.com/open-telemetry/otel-arrow/go/pkg/otel/constants"
	"github.com/open-telemetry/otel-arrow/go/pkg/werror"
)

type EHistogramDataPointBucketsIds struct {
	ID           int
	Offset       int
	BucketCounts int
}

func NewEHistogramDataPointBucketsIds(schema *arrow.Schema, parent string) (*EHistogramDataPointBucketsIds, error) {
	ID, parentDT, err := arrowutils.StructFieldIDFromSchema(schema, parent)
	if err != nil {
		return nil, werror.Wrap(err)
	}
	offset, _ := arrowutils.FieldIDFromStruct(parentDT, constants.ExpHistogramOffset)
	bucketCounts, _ := arrowutils.FieldIDFromStruct(parentDT, constants.ExpHistogramBucketCounts)

	if offset == arrowutils.AbsentFieldID && bucketCounts == arrowutils.AbsentFieldID {
		ID = arrowutils.AbsentFieldID
	}

	return &EHistogramDataPointBucketsIds{
		ID:           ID,
		Offset:       offset,
		BucketCounts: bucketCounts,
	}, nil
}

func AppendUnivariateEHistogramDataPointBucketsInto(dpBuckets pmetric.ExponentialHistogramDataPointBuckets, ehdp *array.Struct, ids *EHistogramDataPointBucketsIds, row int) error {
	if ehdp == nil {
		return nil
	}

	offsetArr := ehdp.Field(ids.Offset)
	if offsetArr != nil {
		if i32OffsetArr, ok := offsetArr.(*array.Int32); ok {
			dpBuckets.SetOffset(i32OffsetArr.Value(row))
		} else {
			return werror.Wrap(ErrNotArrayInt32)
		}
	}

	bucketCountsArr := ehdp.Field(ids.BucketCounts)
	if bucketCountsArr != nil {
		if i64BucketCountsArr, ok := bucketCountsArr.(*array.List); ok {
			start := int(i64BucketCountsArr.Offsets()[row])
			end := int(i64BucketCountsArr.Offsets()[row+1])
			values := i64BucketCountsArr.ListValues()

			if v, ok := values.(*array.Uint64); ok {
				dpbs := dpBuckets.BucketCounts()
				dpbs.EnsureCapacity(end - start)
				for i := start; i < end; i++ {
					dpbs.Append(v.Value(i))
				}
			}
		} else {
			return werror.Wrap(ErrNotArrayList)
		}
	}

	return nil
}
