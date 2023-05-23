package fqdn

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// isalnum(3p) in POSIX locale
func isalnum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

const (
	maxHostnameLen = 254
)

// Validate hostname, based on musl-c version of this function.
func isValidHostname(s string) bool {
	if len(s) > maxHostnameLen {
		return false
	}

	for _, c := range s {
		if !(c >= 0x80 || c == '.' || c == '-' || isalnum(c)) {
			return false
		}
	}

	return true
}

func parseHostLine(host string, line string) (string, bool) {
	const (
		StateSkipWhite = iota
		StateIp
		StateCanonFirst
		StateCanon
		StateAliasFirst
		StateAlias
	)

	var (
		canon     string
		state     int
		nextState int

		i     int
		start int
	)

	isWhite := func(b byte) bool {
		return b == ' ' || b == '\t'
	}
	isLast := func() bool {
		return i == len(line)-1 || isWhite(line[i+1])
	}
	partStr := func() string {
		return line[start : i+1]
	}

	state = StateSkipWhite
	nextState = StateIp

	debug("Looking for %q in %q", host, line)
	for i = 0; i < len(line); i += 1 {
		debug("%03d: character %q, state: %d, nstate: %d",
			i, line[i], state, nextState)

		if line[i] == '#' {
			debug("%03d: found comment, terminating", i)
			break
		}

		switch state {
		case StateSkipWhite:
			if !isWhite(line[i]) {
				state = nextState
				i -= 1
			}
		case StateIp:
			if isLast() {
				state = StateSkipWhite
				nextState = StateCanonFirst
			}
		case StateCanonFirst:
			start = i
			state = StateCanon
			i -= 1
		case StateCanon:
			debug("Canon so far: %q", partStr())
			if isLast() {
				canon = partStr()
				if !isValidHostname(canon) {
					return "", false
				}

				if canon == host {
					debug("Canon match")
					return canon, true
				}

				state = StateSkipWhite
				nextState = StateAliasFirst
			}
		case StateAliasFirst:
			start = i
			state = StateAlias
			i -= 1
		case StateAlias:
			debug("Alias so far: %q", partStr())
			if isLast() {
				alias := partStr()
				if alias == host {
					debug("Alias match")
					return canon, true
				}

				state = StateSkipWhite
				nextState = StateAliasFirst
			}
		default:
			panic(fmt.Sprintf("BUG: State not handled: %d", state))
		}
	}

	debug("No match")
	return "", false
}

// Reads hosts(5) file and tries to get canonical name for host.
func fromHosts(host string) (string, error) {
	var (
		fqdn string
		line string
		err  error
		file *os.File
		r    *bufio.Reader
		ok   bool
	)

	file, err = os.Open(hostsPath)
	if err != nil {
		err = fmt.Errorf("cannot open hosts file: %w", err)
		goto out
	}
	defer file.Close()

	r = bufio.NewReader(file)
	for line, err = readline(r); err == nil; line, err = readline(r) {
		fqdn, ok = parseHostLine(host, line)
		if ok {
			goto out
		}
	}

	if err != io.EOF {
		err = fmt.Errorf("failed to read file: %w", err)
		goto out
	}
	err = errFqdnNotFound{}

out:
	return fqdn, err
}

func fromLookup(host string) (string, error) {
	var (
		fqdn  string
		err   error
		addrs []net.IP
		hosts []string
	)

	fqdn, err = net.LookupCNAME(host)
	if err == nil && len(fqdn) != 0 {
		debug("LookupCNAME success: %q", fqdn)
		goto out
	}
	debug("LookupCNAME failed: %v", err)

	debug("Looking up: %q", host)
	addrs, err = net.LookupIP(host)
	if err != nil {
		err = errFqdnNotFound{err}
		goto out
	}
	debug("Resolved addrs: %q", addrs)

	for _, addr := range addrs {
		debug("Trying: %q", addr)
		hosts, err = net.LookupAddr(addr.String())
		// On windows it can return err == nil but empty list of hosts
		if err != nil || len(hosts) == 0 {
			continue
		}
		debug("Resolved hosts: %q", hosts)

		// First one should be the canonical hostname
		fqdn = hosts[0]

		goto out
	}

	err = errFqdnNotFound{}
out:
	// For some reason we wanted the canonical hostname without
	// trailing dot. So if it is present, strip it.
	if len(fqdn) > 0 && fqdn[len(fqdn)-1] == '.' {
		fqdn = fqdn[:len(fqdn)-1]
	}

	return fqdn, err
}

// Try to get fully qualified hostname for current machine.
//
// It tries to mimic how `hostname -f` works, so except for few edge cases you
// should get the same result from both. One thing that needs to be mentioned is
// that it does not guarantee that you get back fqdn. There is no way to do that
// and `hostname -f` can also return non-fqdn hostname if your /etc/hosts is
// fucked up.
//
// It checks few sources in this order:
//
// 1. hosts file
//	It parses hosts file if present and readable and returns first canonical
//	hostname that also references your hostname. See hosts(5) for more
//	details.
// 2. dns lookup
//	If lookup in hosts file fails, it tries to ask dns.
//
// If none of steps above succeeds, ErrFqdnNotFound is returned as error. You
// will probably want to just use output from os.Hostname() at that point.
func FqdnHostname() (string, error) {
	var (
		fqdn string
		host string
		err  error
	)

	host, err = os.Hostname()
	if err != nil {
		err = errHostnameFailed{err}
		goto out
	}
	debug("Hostname: %q", host)

	fqdn, err = fromHosts(host)
	if err == nil {
		debug("fqdn fetched from hosts: %q", fqdn)
		goto out
	}

	fqdn, err = fromLookup(host)
	if err == nil {
		debug("fqdn fetched from lookup: %q", fqdn)
		goto out
	}

	debug("fqdn fetch failed: %v", err)
out:
	return fqdn, err
}
