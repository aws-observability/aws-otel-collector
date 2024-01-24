package fqdn

import (
	"bufio"
	"io"
)

// Read lines from r. It strips the line terminators and handles case when last
// line is not terminated.
func readline(r *bufio.Reader) (string, error) {
	s, e := r.ReadString('\n')

	if e == io.EOF && len(s) != 0 {
		e = nil
	}

	s = chomp(s)

	return s, e
}
