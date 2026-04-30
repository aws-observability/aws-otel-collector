// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PlaylistDataType Rum replay playlist resource type.
type PlaylistDataType string

// List of PlaylistDataType.
const (
	PLAYLISTDATATYPE_RUM_REPLAY_PLAYLIST PlaylistDataType = "rum_replay_playlist"
)

var allowedPlaylistDataTypeEnumValues = []PlaylistDataType{
	PLAYLISTDATATYPE_RUM_REPLAY_PLAYLIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PlaylistDataType) GetAllowedValues() []PlaylistDataType {
	return allowedPlaylistDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PlaylistDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PlaylistDataType(value)
	return nil
}

// NewPlaylistDataTypeFromValue returns a pointer to a valid PlaylistDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlaylistDataTypeFromValue(v string) (*PlaylistDataType, error) {
	ev := PlaylistDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PlaylistDataType: valid values are %v", v, allowedPlaylistDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlaylistDataType) IsValid() bool {
	for _, existing := range allowedPlaylistDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaylistDataType value.
func (v PlaylistDataType) Ptr() *PlaylistDataType {
	return &v
}
