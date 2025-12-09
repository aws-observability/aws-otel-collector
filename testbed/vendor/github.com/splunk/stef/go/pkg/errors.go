package pkg

import "errors"

var ErrMultimap = errors.New("invalid multimap")
var ErrMultimapCountLimit = errors.New("too many elements in the multimap")
var ErrInvalidRefNum = errors.New("invalid refNum")
var ErrInvalidOneOfType = errors.New("invalid oneof type")

var ErrInvalidHeader = errors.New("invalid FixedHeader")
var ErrInvalidHeaderSignature = errors.New("invalid FixedHeader signature")
var ErrInvalidFormatVersion = errors.New("invalid format version in the FixedHeader")
var ErrInvalidCompression = errors.New("invalid compression method")
