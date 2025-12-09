package mmh3

import (
	"encoding/binary"
	"math"
	"math/bits"
	"reflect"
	"unsafe"
)

const (
	h32c1 uint32 = 0xcc9e2d51
	h32c2 uint32 = 0x1b873593
	h64c1 uint64 = 0x87c37b91114253d5
	h64c2 uint64 = 0x4cf5ad432745937f

	h128BlockSize = 16
)

type Hash128Value struct {
	high, low uint64
}

// Values returns the two 8 byte values that compose the hash
func (h Hash128Value) Values() (uint64, uint64) {
	return h.high, h.low
}

// Write writes the hash to given buffer which should be 16 bytes long
func (h Hash128Value) Write(out []byte) {
	if len(out) != 16 {
		panic("provided return buffer should be 16 bytes")
	}

	rh := *(*reflect.SliceHeader)(unsafe.Pointer(&out))
	rh.Len = 2

	b := *(*[]uint64)(unsafe.Pointer(&rh))
	b[0] = h.high
	b[1] = h.low

	rh.Len = 16
}

// Bytes returns the 16 bytes that compose the hash
func (h Hash128Value) Bytes() []byte {
	b := make([]byte, 16)
	h.Write(b)
	return b
}

func Hash32(key []byte) uint32 {
	length := len(key)
	if length == 0 {
		return 0
	}
	nblocks := length / 4
	var h, k uint32
	for i := 0; i < nblocks; i++ {
		k = binary.LittleEndian.Uint32(key[i*4:])
		k *= h32c1
		k = (k << 15) | (k >> (32 - 15))
		k *= h32c2
		h ^= k
		h = (h << 13) | (h >> (32 - 13))
		h = (h * 5) + 0xe6546b64
	}
	k = 0
	tailIndex := nblocks * 4
	switch length & 3 {
	case 3:
		k ^= uint32(key[tailIndex+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[tailIndex+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[tailIndex])
		k *= h32c1
		k = (k << 15) | (k >> (32 - 15))
		k *= h32c2
		h ^= k
	}
	h ^= uint32(length)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

// Hash128 s a version of MurmurHash which is designed to run only on little-endian processors
func Hash128(key []byte) Hash128Value {
	length := len(key)
	if length == 0 {
		return Hash128Value{}
	}

	nblocks := length / 16
	var h1, h2 uint64

	values := (*(*[]uint64)(unsafe.Pointer(&key)))[: nblocks*2 : nblocks*2]
	for i := 0; i < len(values); i += 2 {
		h1, h2 = merge128(h1, h2, values[i], values[i+1])
	}

	var k1, k2 uint64
	tailIndex := nblocks * 16
	switch length & 15 {
	case 15:
		k2 ^= uint64(key[tailIndex+14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(key[tailIndex+13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(key[tailIndex+12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(key[tailIndex+11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(key[tailIndex+10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(key[tailIndex+9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(key[tailIndex+8])
		k2 *= h64c2
		k2 = bits.RotateLeft64(k2, 33)
		k2 *= h64c1
		h2 ^= k2
		fallthrough
	case 8:
		k1 ^= uint64(key[tailIndex+7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(key[tailIndex+6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(key[tailIndex+5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(key[tailIndex+4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(key[tailIndex+3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(key[tailIndex+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(key[tailIndex+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(key[tailIndex])
		k1 *= h64c1
		k1 = bits.RotateLeft64(k1, 31)
		k1 *= h64c2
		h1 ^= k1
	}
	h1 ^= uint64(length)
	h2 ^= uint64(length)
	h1 += h2
	h2 += h1
	h1 ^= h1 >> 33
	h1 *= 0xff51afd7ed558ccd
	h1 ^= h1 >> 33
	h1 *= 0xc4ceb9fe1a85ec53
	h1 ^= h1 >> 33
	h2 ^= h2 >> 33
	h2 *= 0xff51afd7ed558ccd
	h2 ^= h2 >> 33
	h2 *= 0xc4ceb9fe1a85ec53
	h2 ^= h2 >> 33
	h1 += h2
	h2 += h1

	return Hash128Value{high: h1, low: h2}
}

// Hash128x64 calls WriteHash128x64 with an allocated output buffer
func Hash128x64(key []byte) []byte {
	return Hash128(key).Bytes()
}

// WriteHash128x64 creates a hash of key and writes it to ret
func WriteHash128x64(key, ret []byte) {
	Hash128(key).Write(ret)
}

type HashWriter128 struct {
	h1         uint64
	h2         uint64
	tail       [h128BlockSize]byte
	tailLength int
	length     int
}

func (hw *HashWriter128) Reset() {
	hw.h1 = 0
	hw.h2 = 0
	hw.length = 0
	hw.tailLength = 0
}

func (hw *HashWriter128) Size() int {
	return h128BlockSize
}

func (hw *HashWriter128) BlockSize() int {
	return h128BlockSize
}

// Write writes the given bytes to the hash
func (hw *HashWriter128) Write(b []byte) (int, error) {
	hw.AddBytes(b)
	return len(b), nil
}

// WriteString writes the given string to the hash
func (hw *HashWriter128) WriteString(s string) (int, error) {
	hw.AddString(s)
	return len(s), nil
}

// Sum completes the hash and _appends_ the hash to the given input buffer
// This is kept for backwards compatibility
func (hw *HashWriter128) Sum(b []byte) []byte {
	h1, h2 := hw.Sum128().Values()
	var retbuf [8]byte
	binary.LittleEndian.PutUint64(retbuf[:], h1)
	b = append(b, retbuf[:]...)
	binary.LittleEndian.PutUint64(retbuf[:], h2)
	b = append(b, retbuf[:]...)
	return b
}

// AddBytes adds the given bytes to the hash
func (hw *HashWriter128) AddBytes(input []byte) {
	chunkStart := 0 // The index in to the input where we will begin reading full chunks of data
	inputLength := len(input)

	hw.length += inputLength

	// If there are bytes remaining in the tail, attempt to form
	// a full block and process that first
	if hw.tailLength > 0 {
		bytesNeeded := h128BlockSize - hw.tailLength

		// If the input isn't large enough to form a full block with the tail
		// copy it in to the tail and return
		if bytesNeeded > inputLength {
			copy(hw.tail[hw.tailLength:], input)
			hw.tailLength += inputLength
			return
		}

		copy(hw.tail[hw.tailLength:], input[0:bytesNeeded])
		chunkStart += bytesNeeded
		inputLength -= bytesNeeded

		k1Buf := hw.tail[0:8]
		k2Buf := hw.tail[8:]

		k1 := *(*uint64)(unsafe.Pointer(&k1Buf[0]))
		k2 := *(*uint64)(unsafe.Pointer(&k2Buf[0]))

		hw.h1, hw.h2 = merge128(hw.h1, hw.h2, k1, k2)
	}

	chunks := inputLength / h128BlockSize

	if chunks > 0 {
		chunkInput := input
		if chunkStart > 0 {
			chunkInput = input[chunkStart:]
		}

		values := (*(*[]uint64)(unsafe.Pointer(&chunkInput)))[: chunks*2 : chunks*2]
		for i := 0; i < len(values); i += 2 {
			hw.h1, hw.h2 = merge128(hw.h1, hw.h2, values[i], values[i+1])
		}
	}

	tailLength := inputLength % h128BlockSize

	hw.tailLength = tailLength

	if tailLength == 0 {
		return
	}

	// Copy remaining bytes in to tail
	tailStart := inputLength - tailLength + chunkStart
	copy(hw.tail[:], input[tailStart:])
}

// AddString adds the given string to the hash
func (hw *HashWriter128) AddString(key string) {
	if len(key) == 0 {
		return
	}
	sh := (*reflect.StringHeader)(unsafe.Pointer(&key))
	byteSlice := (*[math.MaxInt32 - 1]byte)(unsafe.Pointer(sh.Data))[:sh.Len:sh.Len]
	hw.AddBytes(byteSlice)
}

// Sum completes the hash value and returns it.  Reset should be called after calling Sum
func (hw *HashWriter128) Sum128() Hash128Value {
	h1, h2 := hw.h1, hw.h2
	k1, k2 := uint64(0), uint64(0)

	switch hw.tailLength {
	case 15:
		k2 ^= uint64(hw.tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(hw.tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(hw.tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(hw.tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(hw.tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(hw.tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(hw.tail[8])
		k2 *= h64c2
		k2 = bits.RotateLeft64(k2, 33)
		k2 *= h64c1
		h2 ^= k2
		fallthrough
	case 8:
		k1 ^= uint64(hw.tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(hw.tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(hw.tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(hw.tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(hw.tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(hw.tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(hw.tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(hw.tail[0])
		k1 *= h64c1
		k1 = bits.RotateLeft64(k1, 31)
		k1 *= h64c2
		h1 ^= k1
	}

	h1 ^= uint64(hw.length)
	h2 ^= uint64(hw.length)
	h1 += h2
	h2 += h1
	h1 ^= h1 >> 33
	h1 *= 0xff51afd7ed558ccd
	h1 ^= h1 >> 33
	h1 *= 0xc4ceb9fe1a85ec53
	h1 ^= h1 >> 33
	h2 ^= h2 >> 33
	h2 *= 0xff51afd7ed558ccd
	h2 ^= h2 >> 33
	h2 *= 0xc4ceb9fe1a85ec53
	h2 ^= h2 >> 33
	h1 += h2
	h2 += h1

	return Hash128Value{high: h1, low: h2}
}

// merge128 merges k1 and k2 with h1 and h2 and returns the updated (h1, h2) values
func merge128(h1, h2, k1, k2 uint64) (uint64, uint64) {
	k1 *= h64c1
	k1 = bits.RotateLeft64(k1, 31)
	k1 *= h64c2

	h1 ^= k1
	h1 = bits.RotateLeft64(h1, 27)
	h1 += h2
	h1 = h1*5 + 0x52dce729

	k2 *= h64c2
	k2 = bits.RotateLeft64(k2, 33)
	k2 *= h64c1

	h2 ^= k2
	h2 = bits.RotateLeft64(h2, 31)
	h2 += h1
	h2 = h2*5 + 0x38495ab5

	return h1, h2
}
