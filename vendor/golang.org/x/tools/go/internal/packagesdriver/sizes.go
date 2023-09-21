// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package packagesdriver fetches type sizes for go/packages and go/analysis.
package packagesdriver

import (
	"context"
	"fmt"
<<<<<<< HEAD
	"go/types"
=======
>>>>>>> main
	"strings"

	"golang.org/x/tools/internal/gocommand"
)

var debug = false

<<<<<<< HEAD
func GetSizesGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (types.Sizes, error) {
=======
func GetSizesForArgsGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (string, string, error) {
>>>>>>> main
	inv.Verb = "list"
	inv.Args = []string{"-f", "{{context.GOARCH}} {{context.Compiler}}", "--", "unsafe"}
	stdout, stderr, friendlyErr, rawErr := gocmdRunner.RunRaw(ctx, inv)
	var goarch, compiler string
	if rawErr != nil {
		if rawErrMsg := rawErr.Error(); strings.Contains(rawErrMsg, "cannot find main module") || strings.Contains(rawErrMsg, "go.mod file not found") {
			// User's running outside of a module. All bets are off. Get GOARCH and guess compiler is gc.
			// TODO(matloob): Is this a problem in practice?
			inv.Verb = "env"
			inv.Args = []string{"GOARCH"}
			envout, enverr := gocmdRunner.Run(ctx, inv)
			if enverr != nil {
<<<<<<< HEAD
				return nil, enverr
=======
				return "", "", enverr
>>>>>>> main
			}
			goarch = strings.TrimSpace(envout.String())
			compiler = "gc"
		} else {
<<<<<<< HEAD
			return nil, friendlyErr
=======
			return "", "", friendlyErr
>>>>>>> main
		}
	} else {
		fields := strings.Fields(stdout.String())
		if len(fields) < 2 {
<<<<<<< HEAD
			return nil, fmt.Errorf("could not parse GOARCH and Go compiler in format \"<GOARCH> <compiler>\":\nstdout: <<%s>>\nstderr: <<%s>>",
=======
			return "", "", fmt.Errorf("could not parse GOARCH and Go compiler in format \"<GOARCH> <compiler>\":\nstdout: <<%s>>\nstderr: <<%s>>",
>>>>>>> main
				stdout.String(), stderr.String())
		}
		goarch = fields[0]
		compiler = fields[1]
	}
<<<<<<< HEAD
	return types.SizesFor(compiler, goarch), nil
=======
	return compiler, goarch, nil
>>>>>>> main
}
