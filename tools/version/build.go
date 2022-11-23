// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"bytes"
	"fmt"
	"log"
)

var (
	// Version variable will be replaced at link time after `make` has been run.
	Version = "latest"
	// GitHash variable will be replaced at link time after `make` has been run.
	GitHash string
	Date    string
)

func init() {
	log.Printf("ADOT Collector version: %v\n", Version)
}

// Info has properties about the build and runtime.
type Info [][2]string

// String returns a formatted string, with linebreaks, intended to be displayed
// on stdout.
func (i Info) String() string {
	buf := new(bytes.Buffer)
	maxRow1Alignment := 0
	for _, prop := range i {
		if cl0 := len(prop[0]); cl0 > maxRow1Alignment {
			maxRow1Alignment = cl0
		}
	}

	for _, prop := range i {
		// Then finally print them with left alignment
		fmt.Fprintf(buf, "%*s %s\n", -maxRow1Alignment, prop[0], prop[1])
	}
	return buf.String()
}
