package process

import (
	"encoding/binary"
	"hash/fnv"
	"math"
	"slices"
	"sort"
	"sync"
)

// V3TagEncoder is an optimized version of V2TagEncoder - it uses a deterministic index for the same set of tags.
// V2 had 2 main issues:
//  - Calling Encode multiple times with the same tags would produce different indexes
//	- Indexes were not deterministic: same tags in different order would produce different indexes
// V3 sacrifices encoding performance for smaller buffer size and faster decoding.
// V3 also guarantees the decoded tags are always sorted, which is useful for many applications.
// V3 preserves V2's format, hence, we rely on the same decoding logic as V2.
type V3TagEncoder struct {
	// tags are stored as a map of tag name to position in the buffer
	tags map[string]uint32
	// order is the list of tags in the order they were added
	order []string
	// footer is the encoded footer that contains the tag positions
	footer []byte
	// tagPosition is the current position in the buffer where the next tag will be written
	tagPosition uint32

	// tagSetCache is a cache of previously seen tag sets to avoid re-encoding the same set
	// The key is a hash of the sorted tags, and the value is the index in the footer
	// This allows us to return the same index for the same set of tags.
	tagSetCache map[uint64]int

	// internal buffer for reuse
	bufInt32 [4]byte
	bufInt16 [2]byte
}

var v3TagSetCachePool = sync.Pool{
	New: func() interface{} {
		cache := make(map[uint64]int)
		return &cache
	},
}

// NewV3TagEncoder creates a new V3TagEncoder with pre-allocated buffers and pools for reusing memory to reduce allocations.
func NewV3TagEncoder() TagEncoder {
	footer := *footerPool.Get().(*[]byte)
	order := *orderPool.Get().(*[]string)
	tags := *tagsPool.Get().(*map[string]uint32)
	cache := *v3TagSetCachePool.Get().(*map[uint64]int)

	return &V3TagEncoder{
		tags:        tags,
		order:       order[:0],
		footer:      footer[:0],
		tagPosition: v2PreambleLength,
		tagSetCache: cache,
		bufInt32:    [4]byte{},
		bufInt16:    [2]byte{},
	}
}

// Buffer returns the encoded buffer with the tags and footer.
// The buffer format is:
// - 1 byte version (3 for V3)
// - 4 bytes footer position (where the footer starts in the buffer)
// - N tags, each with a 2-byte length followed by the tag bytes
// - Footer: 2 bytes tag count, followed by 4 bytes per tag (positions
//   in the buffer where the tag starts)
func (t *V3TagEncoder) Buffer() []byte {
	buffer := make([]byte, int(t.tagPosition)+len(t.footer))
	pos := 0

	// Write version
	buffer[pos] = version3
	pos++

	// Write footer position
	binary.LittleEndian.PutUint32(buffer[pos:], t.tagPosition)
	pos += lenUint32

	// Write tags - each tag is prefixed with its length
	// so while decoding we can read the tag length first and know exactly how many bytes to read for the tag.
	// It allows more efficient decoding.
	for _, tag := range t.order {
		binary.LittleEndian.PutUint16(buffer[pos:], uint16(len(tag)))
		pos += lenUint16
		copy(buffer[pos:], tag)
		pos += len(tag)
	}

	// Write footer
	copy(buffer[pos:], t.footer)

	// Return pooled objects
	footerPool.Put(&t.footer)
	orderPool.Put(&t.order)

	clear(t.tags)
	tagsPool.Put(&t.tags)

	clear(t.tagSetCache)
	v3TagSetCachePool.Put(&t.tagSetCache)

	return buffer
}

// Encode encodes a slice of tags into the V3 format.
// It returns the index in the buffer where the footer for this set of tags starts.
// It modifies the input slice in place to deduplicate and sort the tags.
func (t *V3TagEncoder) Encode(tags []string) int {
	if len(tags) == 0 {
		return -1
	}

	// Sort first to enable deduplication without seen map
	sort.Strings(tags)

	if len(tags) > math.MaxUint16 {
		tags = tags[0:math.MaxUint16]
	}

	// Deduplicate in place after sorting
	n := 0
	for i, tag := range tags {
		if len(tag) > math.MaxUint16 {
			tags[i] = tag[0:math.MaxUint16]
		}
		if i == 0 || tags[i-1] != tag {
			tags[n] = tags[i]
			n++
		}
	}
	tags = tags[:n]

	// Assign indices for tags and calculate hash in one loop
	hasher := fnv.New64a()
	for _, tag := range tags {
		tagPos, exists := t.tags[tag]
		if !exists {
			// New tag - add to dictionary
			tagPos = t.tagPosition
			t.tags[tag] = tagPos
			t.order = append(t.order, tag)
			t.tagPosition += lenUint16 + uint32(len(tag)) // 2 bytes length + tag data
		}
		// Add tag index to hash
		binary.LittleEndian.PutUint32(t.bufInt32[:], tagPos)
		_, _ = hasher.Write(t.bufInt32[:])
	}
	hash := hasher.Sum64()

	// Check if we've seen this exact tag set before (deterministic)
	if existingIndex, exists := t.tagSetCache[hash]; exists {
		return existingIndex
	}

	// Build footer entry with sorted tags (like V2, but sorted)
	footerIndex := len(t.footer)

	// Pre-grow footer to avoid multiple reallocations
	// We need: 2 bytes (tag count) + 4 bytes per tag (positions)
	footerSize := lenUint16 + lenUint32*len(tags)
	t.footer = slices.Grow(t.footer, footerSize)[:len(t.footer)+footerSize]
	pos := footerIndex

	// Write number of tags
	binary.LittleEndian.PutUint16(t.footer[pos:], uint16(len(tags)))
	pos += lenUint16

	// Write tag positions (recalculate from map - guaranteed to exist)
	for _, tag := range tags {
		tagPos := t.tags[tag]
		binary.LittleEndian.PutUint32(t.footer[pos:], tagPos)
		pos += lenUint32
	}

	// Cache this tag set for deterministic behavior
	t.tagSetCache[hash] = footerIndex

	return footerIndex
}
