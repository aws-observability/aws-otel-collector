package process

import (
	"fmt"
)

func (m *CollectorConnections) GetHostTags(host *Host) []string {
	return m.GetTags(int(host.TagIndex))
}

func (m *CollectorConnections) IterateHostTags(host *Host, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(host.TagIndex), cb)
}

func (m *CollectorConnections) GetResourceTags(resource *ResourceMetadata) []string {
	return m.GetTags(int(resource.TagIndex))
}

func (m *CollectorConnections) IterateResourceTags(resource *ResourceMetadata, cb func(i, total int, tag string) bool) {
	iterateTags(m.EncodedTags, int(resource.TagIndex), cb)
}

func (m *CollectorConnections) GetTags(tagIndex int) []string {
	return getTags(m.EncodedTags, tagIndex)
}

func (m *CollectorConnections) UnsafeIterateTags(tagIndex int, cb func(i, total int, tag []byte) bool) {
	unsafeIterateTags(m.EncodedTags, tagIndex, cb)
}

// GetDNS returns the DNS entries for the given addr.
// The first argument returned is the first DNS entry followed by any additional name resolutions.  Most IPs will
// have a single resolution so this dual format allows us to avoid allocations for the common case.  If there are
// multiple name resolutions, there is no implied priority between the dual values
func (m *CollectorConnections) GetDNS(addr *Addr) (string, []string, error) {
	if m.EncodedDNS != nil {
		return GetDNS(m.EncodedDNS, addr.Ip)
	}
	if m.EncodedDnsLookups != nil && m.EncodedDomainDatabase != nil {
		first, offsets, err := GetDNSV2(m.EncodedDnsLookups, addr.Ip)
		if err != nil {
			return "", nil, err
		}
		firstString, err := getDNSNameFromListByOffset(m.EncodedDomainDatabase, int(first))
		if err != nil {
			return "", nil, err
		}
		var strings []string
		if offsets != nil && (len(offsets) > 0) {
			strings = make([]string, len(offsets))
			for _, off := range offsets {
				s, err := getDNSNameFromListByOffset(m.EncodedDomainDatabase, int(off))
				if err != nil {
					return "", nil, err
				}
				strings = append(strings, s)

			}
		}
		return firstString, strings, nil
	}
	return "", nil, fmt.Errorf("No DNS encoded information")
}

// IterateDNS iterates over all the DNS entries for the given addr, invoking the provided callback for each one
func (m *CollectorConnections) IterateDNS(addr *Addr, cb func(i, total int, entry string) bool) error {
	if m.EncodedDNS != nil {
		return IterateDNS(m.EncodedDNS, addr.Ip, cb)
	}
	if m.EncodedDnsLookups != nil && m.EncodedDomainDatabase != nil {
		var iterError error
		err := IterateDNSV2(m.EncodedDnsLookups, addr.Ip, func(i, total int, offset int32) bool {
			s, err := getDNSNameFromListByOffset(m.EncodedDomainDatabase, int(offset))
			if err == nil {
				return cb(i, total, s)
			}
			iterError = err
			return false
		})
		if err != nil {
			return err
		}
		if iterError != nil {
			return iterError
		}
	}
	return nil
}

// UnsafeIterateDNS iterates over all the DNS entries for the given addr, invoking the provided callback for each one
// The entry returned is only valid for the lifetime of the fields in this message
func (m *CollectorConnections) UnsafeIterateDNS(addr *Addr, cb func(i, total int, entry []byte) bool) error {
	if m.EncodedDNS != nil {
		return UnsafeIterateDNS(m.EncodedDNS, addr.Ip, cb)
	}
	if m.EncodedDnsLookups != nil && m.EncodedDomainDatabase != nil {
		var iterError error
		err := UnsafeIterateDNSV2(m.EncodedDnsLookups, addr.Ip, func(i, total int, offset int32) bool {
			b, err := getDNSNameAsByteSliceByOffset(m.EncodedDomainDatabase, int(offset))
			if err == nil {
				return cb(i, total, b)
			}
			iterError = err
			return false
		})
		if err != nil {
			return err
		}
		if iterError != nil {
			return iterError
		}
	}
	return nil
}

// GetDNSNames returns all the DNS entries
func (m *CollectorConnections) GetDNSNames() ([]string, error) {
	if m.EncodedDNS != nil {
		return getDNSNames(m.EncodedDNS)
	} else if m.EncodedDomainDatabase != nil {
		return getDNSNameListV2(m.EncodedDomainDatabase), nil
	}
	return nil, fmt.Errorf("unknown dns names database")
}

// GetDNSNameByOffset gets the dns name at a given offset
func (m *CollectorConnections) GetDNSNameByOffset(off int32) (string, error) {
	if m.EncodedDomainDatabase == nil {
		return "", fmt.Errorf("no domain database")
	}
	return getDNSNameFromListByOffset(m.EncodedDomainDatabase, int(off))
}

// GetConnectionsTags get tags for a connection
func (m *CollectorConnections) GetConnectionsTags(tagIndex int32) []string {
	return getTags(m.EncodedConnectionsTags, int(tagIndex))
}

// UnsafeIterateConnectionTags iterates the connection tags at the given index, invoking the callback function
// for each one.  The tag slice provided to the callback buffer is unsafe and will not be valid past the end
// of the callback function
func (m *CollectorConnections) UnsafeIterateConnectionTags(tagIndex int32, cb func(i, total int, tag []byte) bool) {
	unsafeIterateTags(m.EncodedConnectionsTags, int(tagIndex), cb)
}
