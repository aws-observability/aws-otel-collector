package pkg

type Compression int

const (
	CompressionNone Compression = 0
	CompressionZstd Compression = 1
	CompressionMask             = 0b11
)
