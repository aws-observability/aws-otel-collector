// +build !windows

package fqdn

const hostnameBin = "hostname"

var hostnameArgs = []string{"-f"} //nolint:gochecknoglobals
