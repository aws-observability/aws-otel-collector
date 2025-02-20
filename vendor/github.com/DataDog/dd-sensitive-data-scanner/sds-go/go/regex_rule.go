package dd_sds

/*
#include <stdlib.h>
#include <dd_sds.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"
	"fmt"
)

type RegexRuleConfig struct {
	Id                 string                   `json:"id"`
	Pattern            string                   `json:"pattern"`
	MatchAction        MatchAction              `json:"match_action"`
	ProximityKeywords  *ProximityKeywordsConfig `json:"proximity_keywords,omitempty"`
	SecondaryValidator SecondaryValidator       `json:"validator,omitempty"`
}

type MatchActionType string
type ReplacementType string

const (
	MatchActionNone          = MatchActionType("None")
	MatchActionRedact        = MatchActionType("Redact")
	MatchActionHash          = MatchActionType("Hash")
	MatchActionPartialRedact = MatchActionType("PartialRedact")

	ReplacementTypeNone         = ReplacementType("none")
	ReplacementTypePlaceholder  = ReplacementType("placeholder")
	ReplacementTypeHash         = ReplacementType("hash")
	ReplacementTypePartialStart = ReplacementType("partial_beginning")
	ReplacementTypePartialEnd   = ReplacementType("partial_end")
)

type SecondaryValidator string

const (
	LuhnChecksum      = SecondaryValidator("LuhnChecksum")
	ChineseIdChecksum = SecondaryValidator("ChineseIdChecksum")
)

type PartialRedactionDirection string

const (
	FirstCharacters = PartialRedactionDirection("FirstCharacters")
	LastCharacters  = PartialRedactionDirection("LastCharacters")
)

// ExtraConfig is used to provide more configuration while creating the rules.
type ExtraConfig struct {
	ProximityKeywords  *ProximityKeywordsConfig
	SecondaryValidator SecondaryValidator
}

// CreateProximityKeywordsConfig creates a ProximityKeywordsConfig.
func CreateProximityKeywordsConfig(lookAheadCharaceterCount uint32, includedKeywords []string, excludedKeywords []string) *ProximityKeywordsConfig {
	if includedKeywords == nil {
		includedKeywords = []string{}
	}
	if excludedKeywords == nil {
		excludedKeywords = []string{}
	}
	return &ProximityKeywordsConfig{
		LookAheadCharacterCount: lookAheadCharaceterCount,
		IncludedKeywords:        includedKeywords,
		ExcludedKeywords:        excludedKeywords,
	}
}

// ProximityKeywordsConfig represents the proximity keyword matching
// for the core library.
type ProximityKeywordsConfig struct {
	LookAheadCharacterCount uint32   `json:"look_ahead_character_count"`
	IncludedKeywords        []string `json:"included_keywords"`
	ExcludedKeywords        []string `json:"excluded_keywords"`
}

// RuleMatch stores the matches reported by the core library.
type RuleMatch struct {
	RuleIdx           uint32
	Path              string
	ReplacementType   ReplacementType
	StartIndex        uint32
	EndIndexExclusive uint32
	ShiftOffset       int32
}

// MatchAction is used to configure the rules.
type MatchAction struct {
	Type MatchActionType
	// used when Type == MatchActionRedact, empty otherwise
	RedactionValue string
	// used when Type == MatchActionPartialRedact, empty otherwise
	CharacterCount uint32
	// used when Type == MatchActionPartialRedact, empty otherwise
	Direction PartialRedactionDirection
}

// NewMatchingRule returns a matching rule with no match _action_.
func NewMatchingRule(id string, pattern string, extraConfig ExtraConfig) RegexRuleConfig {
	return RegexRuleConfig{
		Id:      id,
		Pattern: pattern,
		MatchAction: MatchAction{
			Type: MatchActionNone,
		},
		ProximityKeywords:  extraConfig.ProximityKeywords,
		SecondaryValidator: extraConfig.SecondaryValidator,
	}
}

func (c RegexRuleConfig) CreateRule() (*Rule, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	cdata := C.CString(string(data)) // this call adds the 0, memory has to be freed
	defer C.free(unsafe.Pointer(cdata))

	ptr := C.create_regex_rule(cdata)
	if ptr < 0 {
		return nil, fmt.Errorf("Failed to create regex rule")
	} else {
		return &Rule{nativeRulePtr: int64(ptr)}, nil
	}
}

// NewRedactingRule returns a matching rule redacting events.
func NewRedactingRule(id string, pattern string, redactionValue string, extraConfig ExtraConfig) RegexRuleConfig {
	return RegexRuleConfig{
		Id:      id,
		Pattern: pattern,
		MatchAction: MatchAction{
			Type:           MatchActionRedact,
			RedactionValue: redactionValue,
		},
		ProximityKeywords:  extraConfig.ProximityKeywords,
		SecondaryValidator: extraConfig.SecondaryValidator,
	}
}

// NewHashRule returns a matching rule redacting with hashes.
func NewHashRule(id string, pattern string, extraConfig ExtraConfig) RegexRuleConfig {
	return RegexRuleConfig{
		Id:      id,
		Pattern: pattern,
		MatchAction: MatchAction{
			Type: MatchActionHash,
		},
		ProximityKeywords:  extraConfig.ProximityKeywords,
		SecondaryValidator: extraConfig.SecondaryValidator,
	}
}

// NewPartialRedactRule returns a matching rule partially redacting matches.
func NewPartialRedactRule(id string, pattern string, characterCount uint32, direction PartialRedactionDirection, extraConfig ExtraConfig) RegexRuleConfig {
	return RegexRuleConfig{
		Id:      id,
		Pattern: pattern,
		MatchAction: MatchAction{
			Type:           MatchActionPartialRedact,
			CharacterCount: characterCount,
			Direction:      direction,
		},
		ProximityKeywords:  extraConfig.ProximityKeywords,
		SecondaryValidator: extraConfig.SecondaryValidator,
	}
}

// MarshalJSON marshales the SecondaryValidator.
func (s SecondaryValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"type": string(s),
	})
}

// MarshalJSON marshals the MatchAction in a format understood by the serde rust
// JSON library.
func (m MatchAction) MarshalJSON() ([]byte, error) {
	o := map[string]interface{}{
		"type":         string(m.Type), // serde (rust) will use this field to know what to use for the enum
		"match_action": string(m.Type),
	}

	switch m.Type {
	case MatchActionRedact:
		o["replacement"] = m.RedactionValue
	case MatchActionPartialRedact:
		o["character_count"] = m.CharacterCount
		o["direction"] = string(m.Direction)
	}

	return json.Marshal(o)
}
