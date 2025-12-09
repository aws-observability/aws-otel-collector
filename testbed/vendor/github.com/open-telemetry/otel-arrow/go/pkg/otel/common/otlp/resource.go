/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otlp

import (
	"github.com/apache/arrow-go/v18/arrow"
	"go.opentelemetry.io/collector/pdata/pcommon"

	arrowutils "github.com/open-telemetry/otel-arrow/go/pkg/arrow"
	"github.com/open-telemetry/otel-arrow/go/pkg/otel/constants"
	"github.com/open-telemetry/otel-arrow/go/pkg/werror"
)

type ResourceIds struct {
	Resource               int
	ID                     int
	DroppedAttributesCount int
	SchemaUrl              int
}

func NewResourceIdsFromSchema(schema *arrow.Schema) (*ResourceIds, error) {
	resource, resDT, err := arrowutils.StructFieldIDFromSchema(schema, constants.Resource)
	if err != nil {
		return nil, werror.Wrap(err)
	}

	ID, _ := arrowutils.FieldIDFromStruct(resDT, constants.ID)
	droppedAttributesCount, _ := arrowutils.FieldIDFromStruct(resDT, constants.DroppedAttributesCount)
	schemaUrl, _ := arrowutils.FieldIDFromStruct(resDT, constants.SchemaUrl)

	return &ResourceIds{
		Resource:               resource,
		ID:                     ID,
		DroppedAttributesCount: droppedAttributesCount,
		SchemaUrl:              schemaUrl,
	}, nil
}

func UpdateResourceFromRecord(r pcommon.Resource, record arrow.Record, row int, resIds *ResourceIds, attrsStore *AttributesStore[uint16]) (schemaUrl string, err error) {
	resArr, err := arrowutils.StructFromRecord(record, resIds.Resource, row)
	if err != nil {
		return "", werror.WrapWithContext(err, map[string]interface{}{"row": row})
	}

	// Read schema url
	schemaUrl, err = arrowutils.StringFromStruct(resArr, row, resIds.SchemaUrl)
	if err != nil {
		return "", werror.WrapWithContext(err, map[string]interface{}{"row": row})
	}

	// Read dropped attributes count
	droppedAttributesCount, err := arrowutils.U32FromStruct(resArr, row, resIds.DroppedAttributesCount)
	if err != nil {
		return "", werror.WrapWithContext(err, map[string]interface{}{"row": row})
	}
	r.SetDroppedAttributesCount(droppedAttributesCount)

	// Read attributes
	ID, err := arrowutils.NullableU16FromStruct(resArr, row, resIds.ID)
	if err != nil {
		return "", werror.WrapWithContext(err, map[string]interface{}{"row": row})
	}
	if ID != nil {
		attrs := attrsStore.AttributesByDeltaID(*ID)
		if attrs != nil {
			attrs.CopyTo(r.Attributes())
		}
	}
	return
}

func ResourceIDFromRecord(record arrow.Record, row int, resIDs *ResourceIds) (uint16, error) {
	resStruct, err := arrowutils.StructFromRecord(record, resIDs.Resource, row)
	if err != nil {
		return 0, err
	}
	return arrowutils.U16FromStruct(resStruct, row, resIDs.ID)
}
