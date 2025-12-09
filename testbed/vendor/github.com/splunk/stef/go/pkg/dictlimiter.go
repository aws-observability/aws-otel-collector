package pkg

// SizeLimiter tracks the current (approximate) byte size of dictionaries
// and of of the frame against the specified size limit.
type SizeLimiter struct {
	dictByteSize         uint
	dictByteSizeLimit    uint
	dictSizeLimitReached bool

	frameBitSize      uint
	frameBitSizeLimit uint
}

// Init prepares limiter operation.
// dictByteSizeLimit specifies the limit of size in bytes. 0 means no limit,
// so the size won't be tracked at all.
func (d *SizeLimiter) Init(opts *WriterOptions) {
	d.dictByteSize = 0
	d.frameBitSize = 0
	d.dictByteSizeLimit = opts.MaxTotalDictSize
	d.frameBitSizeLimit = opts.MaxUncompressedFrameByteSize * 8
}

// AddDictElemSize accounts for adding an element of specified bytes size to the dictionary.
func (d *SizeLimiter) AddDictElemSize(elemByteSize uint) {
	if d.dictByteSizeLimit != 0 {
		d.dictByteSize += elemByteSize
		if d.dictByteSize >= d.dictByteSizeLimit {
			d.dictSizeLimitReached = true
		}
	}
}

// AddFrameBits accounts for adding bytes to the frame buffer.
func (d *SizeLimiter) AddFrameBits(bitCount uint) {
	d.frameBitSize += bitCount
}

func (d *SizeLimiter) AddFrameBytes(byteCount uint) {
	d.AddFrameBits(byteCount * 8)
}

// DictLimitReached returns true if the accumulated added element sizes reaches the
// previously defined limit. If specified limit was 0 the limit is never reached
// and this function will always return false.
func (d *SizeLimiter) DictLimitReached() bool {
	return d.dictSizeLimitReached
}

// FrameLimitReached returns true if the accumulated added byte sizes reaches the
// previously defined limit. If specified limit was 0 the limit is never reached
// and this function will always return false.
func (d *SizeLimiter) FrameLimitReached() bool {
	return d.frameBitSizeLimit != 0 && d.frameBitSize >= d.frameBitSizeLimit
}

// ResetDict resets accumulated sizes to 0 and DictLimitReached indicator to false.
// Normally used in conjunction with resetting the dictionary itself.
func (d *SizeLimiter) ResetDict() {
	d.dictByteSize = 0
	d.dictSizeLimitReached = false
}

// ResetFrameSize resets accumulated sizes to 0 and FrameLimitReached indicator to false.
// Normally used after restarting the frame.
func (d *SizeLimiter) ResetFrameSize() {
	d.frameBitSize = 0
}
