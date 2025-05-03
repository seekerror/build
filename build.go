// Package build contains utilities for build-time information.
package build

import (
	"fmt"
	"time"
)

var (
	// Date is the build date in the YYYYMMDD format. May be defined by
	// a linker argument, but effectively defaults to `date +%Y%m%d`.
	Date = time.Now().Format("20060102")
	// GitHash is to be defined by a linker argument as the result of
	// "git rev-parse --short HEAD".
	GitHash = "nohash"
)

// Version holds the full build version, such as "1.4.2.120.4f54bc3d".
type Version struct {
	// Major is the major version.
	Major int
	// Minor is the minor version.
	Minor int
	// Micro is the Micro version.
	Micro int
	// Date is the build date in the YYYYMMDD format.
	Date string
	// Hash is the short git commit hash, if defined, to avoid version
	// collisions for branched or privately built binaries.
	Hash string
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%s.%s", v.Major, v.Minor, v.Micro, v.Date, v.Hash)
}

// NewVersion creates a build version using the global Date and GitHash variables,
// which are expected to be defined externally as linker arguments.
func NewVersion(major, minor, micro int) Version {
	return Version{major, minor, micro, Date, GitHash}
}
