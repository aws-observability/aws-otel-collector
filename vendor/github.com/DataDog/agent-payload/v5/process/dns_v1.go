package process

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/DataDog/mmh3"
)

const bufTooShortStr = "dns buffer is too short"

// DNS data is encoded as a very basic bucketed hash table.  There are three blocks, or buffers, of data:
//
//	The "name" block is all of the unique DNS names.  The length of the name is stored as a varint
//	followed by the name itself
//
//	The "bucket" block contains all of the hash buckets.  The format of each bucket is:
//		varint for number of entries in bucket
//		For each entry in the bucket:
//			varint for length of ip
//			ip bytes
//			varint for number of names associated with the ip
//			Each associated name is encoded as a varint which is the position of the actual name string in the name block
//
//	The "position" block is a list of varints, one for each bucket, where each varint is a pointer to the start
//	of the bucket in the bucket block
//
// The overall buffer is encoded as:
//
//	1 byte indicating version
//	2 bytes indicating the number of buckets
//	varint indicating the length of the name buffer.
//	varint indicating the length of the position buffer
//	varint indicating the position of the "middle" (bucketCount / 2) bucket in the position block
//		We will use this to skip the half of the buckets when searching for the target bucket index
//	position block
//	bucket block
//	name block
//
// Notes:
//
//	Using varints saves space at the cost of not having random access to certain sections of data, particularly the
//	bucket position mapping.  This was a deliberate trade off to reduce the size of the payload and thus memory usage
//
//	Varints are also more finicky to deal with in terms of calculating required space ahead of time.  This increases
//	the implementation complexity, or at least the line count, but we reduce allocations & memory usage by
//	pre-sizing the output buffers
//
// This type is not thread safe

type V1DNSEncoder struct {
	BucketFactor float64
	scratch      [binary.MaxVarintLen64]byte // Used for varint encoding
}

type bucketEntry struct {
	keys []string
	size int
}

// 1 byte for version, 2 byte for bucket count
const dns1Version1PreambleLength = 3

// Used for calculating the number of buckets for a given input map.
// Currently the bucket count is calculated as `len(input) * bucketFactor`
const defaultBucketFactor = 0.75

func NewV1DNSEncoder() DNSEncoderV1 {
	return &V1DNSEncoder{
		BucketFactor: defaultBucketFactor,
	}
}

func (e *V1DNSEncoder) EncodeMapped(dns map[string]*DNSDatabaseEntry, indexToOFfset []int32) ([]byte, error) {
	return nil, fmt.Errorf("EncodeMapped not valid in V1")
}
func (e *V1DNSEncoder) EncodeDomainDatabase(names []string) ([]byte, []int32, error) {
	return nil, nil, fmt.Errorf("EncodeDomainDatabase not valid in V1")
}
func (e *V1DNSEncoder) Encode(dns map[string]*DNSEntry) ([]byte, error) {
	if len(dns) == 0 {
		return nil, nil
	}

	bucketCount := getBucketCount(dns, e.BucketFactor)
	buckets := make([]bucketEntry, bucketCount)

	nameBufferLength := 0
	namePositions := make(map[string]int)
	allBucketsEmpty := true

	// We do three things here:
	//	Build up the keys for each bucket
	//	Calculate the size in bytes for each bucket
	//	Calculate the size of the names buffer
	//		The final value of `nameBufferLength` is the size of the name buffer
	for ip, entry := range dns {
		if len(entry.Names) == 0 {
			continue
		}

		allBucketsEmpty = false

		bucket := int(mmh3.Hash32([]byte(ip))) % bucketCount

		buckets[bucket].keys = append(buckets[bucket].keys, ip)

		buckets[bucket].size += e.varIntSize(len(ip))
		buckets[bucket].size += len(ip)
		buckets[bucket].size += e.varIntSize(len(entry.Names))

		for _, name := range entry.Names {
			position, ok := namePositions[name]
			if !ok {
				position = nameBufferLength // Position is at the current end of the name buffer
				namePositions[name] = position

				nameBufferLength += e.varIntSize(len(name))
				nameBufferLength += len(name)
			}

			buckets[bucket].size += e.varIntSize(position)
		}
	}

	// Exit early if all the buckets are empty
	if allBucketsEmpty {
		return nil, nil
	}

	bucketBufferLength := 0
	positionBufferLength := 0

	// We encode the position of the "middle" bucket in the position buffer as an optimization for reads that
	// lets us skip half of the buckets when scanning for the bucket index
	middleBucket := bucketCount / 2
	middleBucketPosition := 0

	// The size of each bucket also includes the length of the number of keys so add that to each bucket size
	// Calculate the size of the position buffer by summing the length of the varints of each bucket position
	// Calculate the size of the bucket buffer by summing the sizes of all the buckets
	for i := range buckets {
		buckets[i].size += e.varIntSize(len(buckets[i].keys))

		if i == middleBucket {
			middleBucketPosition = positionBufferLength
		}

		positionBufferLength += e.varIntSize(bucketBufferLength)

		bucketBufferLength += buckets[i].size
	}

	var bucketCountBuf [2]byte
	binary.LittleEndian.PutUint16(bucketCountBuf[:], uint16(bucketCount))

	sizeOfPositionBufferLength := e.varIntSize(positionBufferLength)
	sizeOfNameBufferLength := e.varIntSize(nameBufferLength)
	sizeOfMiddleBucketPosition := e.varIntSize(middleBucketPosition)
	metaLength := dns1Version1PreambleLength + sizeOfPositionBufferLength + sizeOfNameBufferLength + sizeOfMiddleBucketPosition

	bufferSize := metaLength + positionBufferLength + bucketBufferLength + nameBufferLength
	buffer := make([]byte, bufferSize)

	metaBuffer := buffer[:0]
	positionBuffer := buffer[metaLength:][:0]
	bucketBuffer := buffer[metaLength+positionBufferLength:][:0]
	nameBuffer := buffer[metaLength+positionBufferLength+bucketBufferLength:]

	metaBuffer = append(metaBuffer, dnsVersion1)
	metaBuffer = append(metaBuffer, bucketCountBuf[:]...)
	metaBuffer = e.appendVarInt(metaBuffer, positionBufferLength)
	metaBuffer = e.appendVarInt(metaBuffer, nameBufferLength)
	metaBuffer = e.appendVarInt(metaBuffer, middleBucketPosition)

	for i := range buckets {
		bucketBuffer = e.appendVarInt(bucketBuffer, len(buckets[i].keys))

		for _, ip := range buckets[i].keys {
			entry := dns[ip]

			bucketBuffer = e.appendVarInt(bucketBuffer, len(ip))
			bucketBuffer = append(bucketBuffer, ip...)
			bucketBuffer = e.appendVarInt(bucketBuffer, len(entry.Names))

			for _, name := range entry.Names {
				position := namePositions[name]

				bucketBuffer = e.appendVarInt(bucketBuffer, position)
			}
		}
	}

	// The position of each bucket is the cumulative sum of the sizes of the previous buckets
	positionCounter := 0
	for i := 0; i < bucketCount; i++ {
		bucketPosition := 0
		if i > 0 {
			bucketPosition = buckets[i-1].size
		}

		positionCounter += bucketPosition

		positionBuffer = e.appendVarInt(positionBuffer, positionCounter)
	}

	for name, position := range namePositions {
		bytesWritten := binary.PutUvarint(nameBuffer[position:], uint64(len(name)))
		copy(nameBuffer[position+bytesWritten:], name)
	}

	return buffer, nil
}

func (e *V1DNSEncoder) varIntSize(value int) int {
	return binary.PutUvarint(e.scratch[0:], uint64(value))
}

func (e *V1DNSEncoder) appendVarInt(buf []byte, value int) []byte {
	bytesWritten := binary.PutUvarint(e.scratch[0:], uint64(value))

	return append(buf, e.scratch[0:bytesWritten]...)
}

func getV1(buf []byte, ip string) (string, []string) {
	var first string
	var names []string

	iterateDNSV1(buf, ip, func(i, total int, entry string) bool {
		if i == 0 {
			first = entry
			if total > 1 {
				names = make([]string, 0, total-1)
			}
		} else {
			names = append(names, entry)
		}
		return true
	})

	return first, names
}

func getDNSNamesV1(buf []byte) []string {
	var names []string
	// skip the preamble
	index := dns1Version1PreambleLength

	_, bytesRead := binary.Uvarint(buf[index:])
	index += bytesRead
	nameBufferLen, bytesRead := binary.Uvarint(buf[index:])

	start := len(buf) - int(nameBufferLen)
	nameBuffer := buf[start:]

	for namePosition := 0; namePosition < len(nameBuffer); {
		nameLength, bytesReadForName := binary.Uvarint(nameBuffer[namePosition:])
		namePosition += bytesReadForName
		name := string(nameBuffer[namePosition : namePosition+int(nameLength)])
		names = append(names, name)
		namePosition += int(nameLength)
	}
	return names
}

func iterateDNSV1(buf []byte, ip string, cb func(i, total int, entry string) bool) error {
	return unsafeIterateDNSV1(buf, ip, func(i, total int, entry []byte) bool {
		return cb(i, total, string(entry))
	})
}

func unsafeIterateDNSV1(buf []byte, ip string, cb func(i, total int, entry []byte) bool) error {
	bufLen := len(buf)

	// Needs 3 bytes so that buf[1:] can convert to uint16
	if bufLen <= 2 {
		return fmt.Errorf(bufTooShortStr)
	}
	// Read overview:
	//	Compute the target bucket for the given ip
	//	Iterate over all the buckets to find position of the given bucket
	// 	Advance to the position of the bucket
	//	For each entry in the bucket:
	//		Compare the key to the given IP and store the comparison result
	//		Iterate through the name positions associated with the key.
	//			If the key was a match, load the name value and add it to the result list.  Return once all names are processed
	//			Otherwise iterate through the name positions to reach the next bucket entry

	bucketCount := int(binary.LittleEndian.Uint16(buf[1:]))

	// skip the preamble
	index := dns1Version1PreambleLength

	if index > bufLen {
		return fmt.Errorf("dns buffer is too short, invalid preamble")
	}
	positionBufferLen, bytesRead := binary.Uvarint(buf[index:])
	index += bytesRead

	if index > bufLen {
		return fmt.Errorf("dns buffer is too short, invalid position buffer length")
	}
	nameBufferLen, bytesRead := binary.Uvarint(buf[index:])
	nameBuffer := buf[len(buf)-int(nameBufferLen):]
	index += bytesRead

	if index > bufLen {
		return fmt.Errorf("dns buffer is too short, invalid middle bucket position")
	}
	middleBucketPosition, bytesRead := binary.Uvarint(buf[index:])
	index += bytesRead

	bucket := int(mmh3.Hash32([]byte(ip))) % bucketCount

	// The length of the metadata is the current read index.  We will use this to calculate the bucket read index below
	metaLength := index

	middleBucket := bucketCount / 2

	startBucket := 0
	endBucket := bucketCount

	if bucket >= middleBucket {
		startBucket = middleBucket
		endBucket = bucketCount

		index += int(middleBucketPosition)
	}

	// Search through the bucket map to find the position of the target bucket
	// Due to varints, we don't know how large the bucket index is
	// We iterate through all the buckets in order to advance the read pointer to the start of the bucket data
	var bucketPosition int

	for i := startBucket; i < endBucket; i++ {
		if index > bufLen {
			return fmt.Errorf("dns buffer is too short, invalid bucket position")
		}
		value, bytesRead := binary.Uvarint(buf[index:])

		index += bytesRead

		if bucket == i {
			bucketPosition = int(value)
			break
		}
	}

	// Move read index to the start of the bucket data.  Skip the metadata and the position buffer
	index = metaLength + int(positionBufferLen) + bucketPosition

	if index > bufLen {
		return fmt.Errorf("dns buffer is too short, invalid bucket length")
	}
	bucketLength, bytesRead := binary.Uvarint(buf[index:])
	index += bytesRead

	for i := 0; i < int(bucketLength); i++ {
		if index > bufLen {
			return fmt.Errorf("dns buffer is too short, invalid key length")
		}
		keyLength, bytesRead := binary.Uvarint(buf[index:])
		index += bytesRead

		if index > bufLen || (index+int(keyLength)) > bufLen {
			return fmt.Errorf("dns buffer is too short, invalid key data`")
		}

		key := buf[index : index+int(keyLength)]
		index += int(keyLength)

		matched := bytes.Equal(key, []byte(ip))

		if index > bufLen {
			return fmt.Errorf("dns buffer is too short, invalid value data`")
		}
		nameCount, bytesRead := binary.Uvarint(buf[index:])
		index += bytesRead

		// Advance through all name positions
		// We still need to do this even if the current entry didn't match in order to get to the next bucket entry
		for j := 0; j < int(nameCount); j++ {
			if index > bufLen {
				return fmt.Errorf("dns buffer is too short, invalid name data`")
			}

			namePosition, bytesRead := binary.Uvarint(buf[index:])
			index += bytesRead

			if !matched {
				continue
			}

			if int(namePosition) > len(nameBuffer) {
				return fmt.Errorf("name buffer is too short, invalid name position`")
			}
			nameLength, bytesReadForName := binary.Uvarint(nameBuffer[int(namePosition):])

			start := int(namePosition) + bytesReadForName

			if start > len(nameBuffer) || start+int(nameLength) > len(nameBuffer) {
				return fmt.Errorf("name buffer is too short, invalid name`")
			}

			if !cb(j, int(nameCount), nameBuffer[start:start+int(nameLength)]) {
				return nil
			}
		}

		if matched {
			return nil
		}
	}

	return nil
}

func getBucketCount(dns map[string]*DNSEntry, bucketFactor float64) int {
	bucketCount := int(float64(len(dns)) * bucketFactor)
	if bucketCount == 0 {
		return 1
	}

	if bucketCount > math.MaxUint16 {
		return math.MaxUint16
	}

	return bucketCount
}

// GetDNS gets the DNS entries for the given IP from the given buffer
func GetDNS(buf []byte, ip string) (string, []string, error) {
	if len(buf) == 0 || ip == "" {
		return "", nil, nil
	}

	switch buf[0] {
	case dnsVersion1:
		first, strings := getV1(buf, ip)
		return first, strings, nil
	}

	return "", nil, fmt.Errorf("Unexpected version %v", buf[0])
}

func getDNSNames(buf []byte) ([]string, error) {
	if len(buf) == 0 {
		return nil, nil
	}

	switch buf[0] {
	case dnsVersion1:
		names := getDNSNamesV1(buf)
		return names, nil
	}
	return nil, fmt.Errorf("Unexpected version %v", buf[0])
}

// IterateDNS invokes the callback function for each DNS entry for the given IP in the given buffer
func IterateDNS(buf []byte, ip string, cb func(i, total int, entry string) bool) error {
	if len(buf) == 0 || ip == "" {
		return nil
	}

	switch buf[0] {
	case dnsVersion1:
		return iterateDNSV1(buf, ip, cb)
	}
	return fmt.Errorf("Unexpected version %v", buf[0])
}

// UnsafeIterateDNS invokes the callback function for each DNS entry for the given IP in the given buffer.
// Each entry is a the slice from the overall buffer.  It should be copied before use
func UnsafeIterateDNS(buf []byte, ip string, cb func(i, total int, entry []byte) bool) error {
	if len(buf) == 0 || ip == "" {
		return nil
	}

	switch buf[0] {
	case dnsVersion1:
		unsafeIterateDNSV1(buf, ip, cb)
		return nil
	}
	return fmt.Errorf("Unexpected version %v", buf[0])
}
