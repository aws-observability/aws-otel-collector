package fqdn

import (
	"net"
	"os"
	"strings"
)

// Get Fully Qualified Domain Name
// returns "unknown" or hostname in case of error
//
// Deprecated:
//             This function has bad API, works poorly and is replace by
//             FqdnHostname. Please please do not use it. It *will* be removed
//             in the next version.
func Get() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".") // return fqdn without trailing dot
		}
	}
	return hostname
}
