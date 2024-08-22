// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package hexes

import (
	"github.com/mdhender/semver"
)

var (
	version = semver.Version{
		Major: 0,
		Minor: 0,
		Patch: 1,
	}
)

func Version() string {
	return version.String()
}
