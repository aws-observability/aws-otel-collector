// +build DEBUG

package fqdn

import "fmt"

func debug(s string, v ...interface{}) {
	if s[len(s)-1] != '\n' {
		s += string('\n')
	}
	fmt.Printf(s, v...)
}
