// +build !windows

package fqdn

func chomp(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	return s
}
