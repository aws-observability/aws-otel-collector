package specs

import "fmt"

const (
	// VersionMajor is for an API incompatible changes
	VersionMajor = 1
	// VersionMinor is for functionality in a backwards-compatible manner
<<<<<<< HEAD
	VersionMinor = 0
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch = 2

	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev = "-dev"
=======
	VersionMinor = 1
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch = 0

	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev = "-rc.3"
>>>>>>> main
)

// Version is the specification version that the package types support.
var Version = fmt.Sprintf("%d.%d.%d%s", VersionMajor, VersionMinor, VersionPatch, VersionDev)
