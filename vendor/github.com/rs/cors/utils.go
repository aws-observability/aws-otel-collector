package cors

<<<<<<< HEAD
import "strings"

const toLower = 'a' - 'A'
=======
import (
	"strings"
)
>>>>>>> main

type converter func(string) string

type wildcard struct {
	prefix string
	suffix string
}

func (w wildcard) match(s string) bool {
	return len(s) >= len(w.prefix)+len(w.suffix) && strings.HasPrefix(s, w.prefix) && strings.HasSuffix(s, w.suffix)
}

<<<<<<< HEAD
// convert converts a list of string using the passed converter function
func convert(s []string, c converter) []string {
	out := []string{}
	for _, i := range s {
		out = append(out, c(i))
=======
// split compounded header values ["foo, bar", "baz"] -> ["foo", "bar", "baz"]
func splitHeaderValues(values []string) []string {
	out := values
	copied := false
	for i, v := range values {
		needsSplit := strings.IndexByte(v, ',') != -1
		if !copied {
			if needsSplit {
				split := strings.Split(v, ",")
				out = make([]string, i, len(values)+len(split)-1)
				copy(out, values[:i])
				for _, s := range split {
					out = append(out, strings.TrimSpace(s))
				}
				copied = true
			}
		} else {
			if needsSplit {
				split := strings.Split(v, ",")
				for _, s := range split {
					out = append(out, strings.TrimSpace(s))
				}
			} else {
				out = append(out, v)
			}
		}
>>>>>>> main
	}
	return out
}

<<<<<<< HEAD
// parseHeaderList tokenize + normalize a string containing a list of headers
func parseHeaderList(headerList string) []string {
	l := len(headerList)
	h := make([]byte, 0, l)
	upper := true
	// Estimate the number headers in order to allocate the right splice size
	t := 0
	for i := 0; i < l; i++ {
		if headerList[i] == ',' {
			t++
		}
	}
	headers := make([]string, 0, t)
	for i := 0; i < l; i++ {
		b := headerList[i]
		switch {
		case b >= 'a' && b <= 'z':
			if upper {
				h = append(h, b-toLower)
			} else {
				h = append(h, b)
			}
		case b >= 'A' && b <= 'Z':
			if !upper {
				h = append(h, b+toLower)
			} else {
				h = append(h, b)
			}
		case b == '-' || b == '_' || b == '.' || (b >= '0' && b <= '9'):
			h = append(h, b)
		}

		if b == ' ' || b == ',' || i == l-1 {
			if len(h) > 0 {
				// Flush the found header
				headers = append(headers, string(h))
				h = h[:0]
				upper = true
			}
		} else {
			upper = b == '-' || b == '_'
		}
	}
	return headers
=======
// convert converts a list of string using the passed converter function
func convert(s []string, c converter) []string {
	out, _ := convertDidCopy(s, c)
	return out
}

// convertDidCopy is same as convert but returns true if it copied the slice
func convertDidCopy(s []string, c converter) ([]string, bool) {
	out := s
	copied := false
	for i, v := range s {
		if !copied {
			v2 := c(v)
			if v2 != v {
				out = make([]string, len(s))
				copy(out, s[:i])
				out[i] = v2
				copied = true
			}
		} else {
			out[i] = c(v)
		}
	}
	return out, copied
>>>>>>> main
}
