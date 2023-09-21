package dns

import "fmt"

// Version is current version of this library.
<<<<<<< HEAD
var Version = v{1, 1, 51}
=======
var Version = v{1, 1, 55}
>>>>>>> main

// v holds the version of this library.
type v struct {
	Major, Minor, Patch int
}

func (v v) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
