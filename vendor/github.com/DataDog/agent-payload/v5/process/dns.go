package process

type DNSEncoderV1 interface {
	Encode(dns map[string]*DNSEntry) ([]byte, error)
}

type DNSEncoderV2 interface {
	// EncodeDomainDatabase encodes a list of domains into a flattened []byte format
	// The names parameter is a list of domain names.
	// The first return value is a []byte containing all the input domains
	// the second argument maps the indexes in the names slice (e.g. 0, 1, 2, 3) to byte offsets into the returned []byte
	EncodeDomainDatabase(names []string) ([]byte, []int32, error)

	// EncodeMapped creates a secondary index from IP address to one or more domains.
	// The dns parameter is a map from IP address to a DNSDatabaseEntry (which is simply a list of indexes into
	// the DNS database which has been previously encoded by EncodeDomainDatabase)
	// The resulting []byte can be read with GetDNSV2, IterateDNSV2, etc.
	EncodeMapped(dns map[string]*DNSDatabaseEntry, indexToOffset []int32) ([]byte, error)
}

const dnsVersion1 byte = 1
const dnsVersion2 byte = 2
