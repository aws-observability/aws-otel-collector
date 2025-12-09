package pkg

import "errors"

type ReadOptions struct {
	// TillEndOfFrame indicates that Read() operation must only succeed if
	// the record is available in the already fetched frame. This ensures
	// that the Reader will not attempt to fetch new frames from the underlying
	// source reader and guarantees that it will not block on I/O.
	// If Read() is attempted when there are no more record remaining in the
	// current frame then ErrEndOfFrame is returned by Read().
	// If TillEndOfFrame is false then Read() will fetch new frames from the
	// underlying reader as needed.
	TillEndOfFrame bool
}

var ErrEndOfFrame = errors.New("end of frame")
