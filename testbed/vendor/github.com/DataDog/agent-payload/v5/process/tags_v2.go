package process

import (
	"encoding/binary"
	"math"
	"sync"
)

// V2TagEncoder operates on the theory that a good portion of the tags for an overall message across connections
// will be duplicated.
//
// Each tag is encoded exactly once in the message at a given position.  Each collection of tags is a list of
// integers representing the position of each tag.  The collection of tag positions is referred to as the footer
// and is appended to the end of the buffer
//
// The format of the buffer is:
//   - 1 byte for meta (currently used for specifying version)
//   - 4 bytes for position of footer blob in overall buffer
//   - N bytes for all tags, stored sequentially.
//     Each tag is 2 bytes for the length of the tag and N bytes for the tag itself
//   - N bytes for the footer blob.  Each entry in the footer is 2 bytes for the number of tags and then N 4 byte
//     integers, each representing the location of the tag in the tag blob
type V2TagEncoder struct {
	tags        map[string]uint32
	order       []string
	footer      []byte
	tagPosition uint32
}

// 1 meta byte + 4 bytes for the index of the footer block
const (
	v2PreambleLength = 1 + 4
	lenUint16        = 2
	lenUint32        = 4
)

var footerPool = sync.Pool{
	New: func() interface{} {
		var footer []byte
		return &footer
	},
}

var orderPool = sync.Pool{
	New: func() interface{} {
		var order []string
		return &order
	},
}

var tagsPool = sync.Pool{
	New: func() interface{} {
		tags := make(map[string]uint32)
		return &tags
	},
}

func NewV2TagEncoder() TagEncoder {
	footer := *footerPool.Get().(*[]byte)
	order := *orderPool.Get().(*[]string)
	tags := *tagsPool.Get().(*map[string]uint32)

	return &V2TagEncoder{
		tags:        tags,
		order:       order[:0],
		footer:      footer[:0],
		tagPosition: v2PreambleLength, // Tags start after the preamble
	}
}

func (t *V2TagEncoder) Buffer() []byte {
	tagsSize := 0

	for _, tag := range t.order {
		tagsSize += 2 + len(tag)
	}

	footerPosition := uint32(v2PreambleLength + tagsSize)

	bufferSize := v2PreambleLength + tagsSize + len(t.footer)
	buffer := make([]byte, 0, bufferSize)
	buffer = append(buffer, version2)

	var intBuf [4]byte
	binary.LittleEndian.PutUint32(intBuf[:], footerPosition)
	buffer = append(buffer, intBuf[:]...)

	var shortBuf [2]byte
	for _, tag := range t.order {
		binary.LittleEndian.PutUint16(shortBuf[:], uint16(len(tag)))
		buffer = append(buffer, shortBuf[:]...)
		buffer = append(buffer, tag...)
	}

	buffer = append(buffer, t.footer...)

	footerPool.Put(&t.footer)
	orderPool.Put(&t.order)

	for k := range t.tags {
		delete(t.tags, k)
	}

	tagsPool.Put(&t.tags)

	return buffer
}

func (t *V2TagEncoder) Encode(tags []string) int {
	if len(tags) == 0 {
		return -1
	}

	var shortBuf [2]byte
	var intBuf [4]byte

	// We only allow 2 bytes for the number of the tags, ensure we don't exceed it
	if len(tags) > math.MaxUint16 {
		tags = tags[0:math.MaxUint16]
	}

	// The index for these tags is the current end of the footer
	tagIndex := len(t.footer)

	binary.LittleEndian.PutUint16(shortBuf[:], uint16(len(tags)))
	t.footer = append(t.footer, shortBuf[:]...)

	for _, tag := range tags {
		// We only allow 2 bytes for the length of the tag, ensure we don't exceed it
		if len(tag) > math.MaxUint16 {
			tag = tag[0:math.MaxUint16]
		}

		position, ok := t.tags[tag]
		if !ok {
			position = t.tagPosition
			t.tagPosition += uint32(2 + len(tag))
			t.tags[tag] = position
			t.order = append(t.order, tag)
		}

		binary.LittleEndian.PutUint32(intBuf[:], position)
		t.footer = append(t.footer, intBuf[:]...)
	}

	return tagIndex
}

func decodeV2(buffer []byte, tagIndex int) []string {
	var tags []string

	iterateV2(buffer, tagIndex, func(i, total int, tag string) bool {
		if i == 0 {
			tags = make([]string, 0, total)
		}
		tags = append(tags, tag)
		return true
	})

	return tags
}

func iterateV2(buffer []byte, tagIndex int, cb func(i, total int, tag string) bool) {
	unsafeIterateV2(buffer, tagIndex, func(i, total int, tag []byte) bool {
		return cb(i, total, string(tag))
	})
}

func unsafeIterateV2(buffer []byte, tagIndex int, cb func(i, total int, tag []byte) bool) {
	if len(buffer) < lenUint32+1 || tagIndex < 0 {
		return
	}
	footerPosition := binary.LittleEndian.Uint32(buffer[1:])

	idx := int(footerPosition) + tagIndex
	if idx >= len(buffer) {
		return
	}
	footerBuffer := buffer[idx:]
	footerIndex := 0

	if len(footerBuffer[footerIndex:]) < lenUint16 {
		return
	}

	numTags := int(binary.LittleEndian.Uint16(footerBuffer[footerIndex:]))
	footerIndex += 2

	for i := 0; i < numTags; i++ {
		if footerIndex >= len(footerBuffer) || len(footerBuffer[footerIndex:]) < lenUint32 {
			continue
		}
		tagPosition := int(binary.LittleEndian.Uint32(footerBuffer[footerIndex:]))

		if tagPosition >= len(buffer) || len(buffer[tagPosition:]) < lenUint16 {
			continue
		}
		tagLength := int(binary.LittleEndian.Uint16(buffer[tagPosition:]))

		start := tagPosition + 2
		end := start + tagLength

		if end > len(buffer) {
			continue
		}

		if !cb(i, numTags, buffer[start:end]) {
			return
		}

		footerIndex += 4
	}
}
