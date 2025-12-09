package metadata

import (
	"fmt"
	"runtime/debug"
)

const packagePath = "github.com/linode/go-metadata"

var (
	Version = "dev"

	// DefaultUserAgent is the default User-Agent sent in HTTP request headers
	DefaultUserAgent string
)

// init attempts to source the version from the build info injected
// at runtime and sets the DefaultUserAgent.
func init() {
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path == packagePath {
				if dep.Replace != nil {
					Version = dep.Replace.Version
				} else {
					Version = dep.Version
				}
				break
			}
		}
	}

	DefaultUserAgent = fmt.Sprintf("go-metadata/%s https://github.com/linode/go-metadata", Version)
}
