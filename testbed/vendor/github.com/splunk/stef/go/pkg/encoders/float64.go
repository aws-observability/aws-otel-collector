package encoders

import (
	"math"
	"math/bits"

	"github.com/splunk/stef/go/pkg"
)

type Float64Encoder struct {
	buf          pkg.BitsWriter
	limiter      *pkg.SizeLimiter
	lastVal      float64
	leadingBits  int
	trailingBits int
}

func (e *Float64Encoder) Init(limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet) error {
	e.limiter = limiter
	return nil
}

func (e *Float64Encoder) IsEqual(val float64) bool {
	return e.lastVal == val
}

func (e *Float64Encoder) Encode(val float64) {
	xorVal := math.Float64bits(val) ^ math.Float64bits(e.lastVal)
	e.lastVal = val
	if xorVal == 0 {
		// Same value.
		e.buf.WriteBit(0)
		e.limiter.AddFrameBits(1)
		return
	}

	leading := bits.LeadingZeros64(xorVal)
	if leading >= 32 {
		// We can encode at most 31 in 5 bits, so clamp to 31.
		// It is OK to assume less leading zeros than in reality. We will just end up storing
		// an extra leading zero bit.
		leading = 31
	}

	trailing := bits.TrailingZeros64(xorVal)

	prevLeading := e.leadingBits
	prevTrailing := e.trailingBits
	sigbits := 64 - leading - trailing

	if prevLeading != -1 && leading >= prevLeading && trailing >= prevTrailing {
		// Fits in previous [leading..trailing] range.
		if 53-prevLeading-prevTrailing < sigbits {
			// Current scheme is smaller than trying reset the range. Use the current scheme.
			e.buf.WriteBits(0b10, 2)
			bitCount := uint(64 - prevLeading - prevTrailing)
			e.buf.WriteBits(xorVal>>prevTrailing, bitCount)
			e.limiter.AddFrameBits(2 + bitCount)
			return
		}
	}

	// Does not fit in previous size, need to store new leading and trailing bit counts.
	e.leadingBits = leading
	e.trailingBits = trailing
	if sigbits == 0 {
		panic("unexpected")
	}

	bitsVal := uint64(0b11)
	bitsVal = (bitsVal << 5) | uint64(leading)
	// Store (sigbits-1) to ensure the value is in 0..63 range to fit in 6 bits.
	bitsVal = (bitsVal << 6) | uint64(sigbits-1)
	e.buf.WriteBits(bitsVal, 13)

	e.buf.WriteBits(xorVal>>trailing, uint(sigbits))

	e.limiter.AddFrameBits(13 + uint(sigbits))
}

func (e *Float64Encoder) CollectColumns(columnSet *pkg.WriteColumnSet) {
	columnSet.SetBits(&e.buf)
}

func (e *Float64Encoder) WriteTo(buf *pkg.BytesWriter) {
	buf.WriteBytes(e.buf.Bytes())
}

func (e *Float64Encoder) Reset() {
	e.lastVal = 0
	e.leadingBits = 0
	e.trailingBits = 0
}

type Float64Decoder struct {
	buf          pkg.BitsReader
	column       *pkg.ReadableColumn
	lastVal      float64
	leadingBits  uint64
	trailingBits uint64
}

func (d *Float64Decoder) Continue() {
	d.buf.Reset(d.column.Data())
}

const (
	// Float64NonIdenticalBit indicates that the value is not identical to the previous value.
	Float64NonIdenticalBit = 0b1000000000000

	// Float64NewLeadingTrailingBit indicates that encoding uses new leading/trailing bit counts.
	Float64NewLeadingTrailingBit = 0b0100000000000

	// Float64LeadingBitMask contains bits that store the leading bit count.
	Float64LeadingBitMask = 0b0011111000000

	// Float64SigBitMask contains bits that store the trailing bit count.
	Float64SigBitMask = 0b0000000111111

	// Float64SigBitsCount the of bits used by Float64SigBitMask, same as the
	// number of bits to shift Float64LeadingBitMask to get its value.
	Float64SigBitsCount = 6
)

func (d *Float64Decoder) Decode(dst *float64) error {
	hdrBits := d.buf.PeekBits(13)
	if hdrBits&Float64NonIdenticalBit == 0 {
		// Value is identical to previous, nothing to do.
		d.buf.Consume(1)
		*dst = d.lastVal
		return nil
	} else {
		var leading, trailing, sigbits uint64
		if hdrBits&Float64NewLeadingTrailingBit == 0 {
			d.buf.Consume(2)

			// Same leading and trailing
			leading = d.leadingBits
			trailing = d.trailingBits
			sigbits = 64 - leading - trailing
		} else {
			d.buf.Consume(13)

			leading = hdrBits & Float64LeadingBitMask >> Float64SigBitsCount
			sigbits = hdrBits & Float64SigBitMask
			sigbits++ // Move from 0..63 to 1..64 range that is the actual value.
			trailing = 64 - leading - sigbits
			d.leadingBits = leading
			d.trailingBits = trailing
		}

		xorVal := d.buf.ReadBits(uint(sigbits))
		xorVal = xorVal << trailing
		d.lastVal = math.Float64frombits(xorVal ^ math.Float64bits(d.lastVal))
	}
	*dst = d.lastVal
	return nil
}

func (d *Float64Decoder) Reset() {
	d.lastVal = 0
	d.leadingBits = 0
	d.trailingBits = 0
}

func (d *Float64Decoder) Init(columns *pkg.ReadColumnSet) error {
	d.column = columns.Column()
	return nil
}
