//go:build linux && !go1.16
// +build linux,!go1.16

/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package userutils

import (
	"github.com/opencontainers/runc/libcontainer/system"
)

// go1.15 and before: use Setgid/Setuid from 3rd party library.
// go1.16 and later: use Setgid/Setuid implemented in go syscall.

func setUid(uid int) (err error) {
	return system.Setuid(uid)
}

func setGid(gid int) (err error) {
	return system.Setgid(gid)
}
