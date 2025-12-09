package pkg

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

// Bytes is a sequence of immutable bytes.
type Bytes string

const EmptyBytes = Bytes("")

func Uint64Compare(left, right uint64) int {
	if left > right {
		return 1
	}
	if left < right {
		return -1
	}
	return 0
}

func Int64Compare(left, right int64) int {
	if left > right {
		return 1
	}
	if left < right {
		return -1
	}
	return 0
}

func BoolCompare(left, right bool) int {
	if left == right {
		return 0
	}
	if left {
		return 1
	}
	return -1
}

func Float64Compare(left, right float64) int {
	if left > right {
		return 1
	}
	if left < right {
		return -1
	}
	return 0
}

func StringCompare(left, right string) int {
	return strings.Compare(left, right)
}

func BytesCompare(left, right Bytes) int {
	return strings.Compare(string(left), string(right))
}

func Uint64Equal(left, right uint64) bool {
	return left == right
}

func Int64Equal(left, right int64) bool {
	return left == right
}

func BoolEqual(left, right bool) bool {
	return left == right
}

func Float64Equal(left, right float64) bool {
	return left == right
}

func StringEqual(left, right string) bool {
	return left == right
}

func BytesEqual(left, right Bytes) bool {
	return left == right
}

func Uint64Random(random *rand.Rand) uint64 {
	return random.Uint64()
}

func Int64Random(random *rand.Rand) int64 {
	return random.Int64()
}

func BoolRandom(random *rand.Rand) bool {
	return random.IntN(2) == 0
}

func Float64Random(random *rand.Rand) float64 {
	return random.Float64()
}

func StringRandom(random *rand.Rand) string {
	return strconv.Itoa(random.IntN(10))
}

func BytesRandom(random *rand.Rand) Bytes {
	return Bytes(StringRandom(random))
}
