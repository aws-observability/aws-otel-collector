/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package common

import (
	"math/rand"

	"github.com/apache/arrow-go/v18/arrow"

	"github.com/open-telemetry/otel-arrow/go/pkg/record_message"
)

func MixUpArrowRecords(rng *rand.Rand, record arrow.Record, relatedRecords []*record_message.RecordMessage) (bool, arrow.Record, []*record_message.RecordMessage) {
	mainRecordChanged := false

	if rng.Intn(100)%2 == 0 {
		// exchange one of the related records with the main record
		relatedRecordPos := rng.Intn(len(relatedRecords))
		relatedRecord := relatedRecords[relatedRecordPos].Record()
		relatedRecords[relatedRecordPos].SetRecord(record)
		record = relatedRecord
		mainRecordChanged = true
	}

	// mix up the related records
	payloadTypes := make([]record_message.PayloadType, len(relatedRecords))
	for i := 0; i < len(relatedRecords); i++ {
		payloadTypes[i] = relatedRecords[i].PayloadType()
	}
	rng.Shuffle(len(payloadTypes), func(i, j int) { payloadTypes[i], payloadTypes[j] = payloadTypes[j], payloadTypes[i] })
	for i := 0; i < len(relatedRecords); i++ {
		relatedRecords[i].SetPayloadType(payloadTypes[i])
	}

	return mainRecordChanged, record, relatedRecords
}
