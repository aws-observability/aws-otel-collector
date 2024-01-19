// +build !DEBUG

package fqdn

// Internal debug functions which by default does nothing. This allows compiler
// to optimize it out so it has no performance impact. If you want the output,
// recompile with `-tags DEBUG`.
func debug(s string, v ...interface{}) {}
