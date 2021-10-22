//go:build linux
// +build linux

/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/opencontainers/runc/libcontainer/user"
	"golang.org/x/sys/unix"
)

var aocUserName = "aoc"
var defaultInstallPath = "/opt/aws/aws-otel-collector/"

type ChownFunc func(name string, uid, gid int) error

var chown ChownFunc = os.Chown

// switchUser swithes the current process to new user
func switchUser(execUser *user.ExecUser) error {
	if err := unix.Setgroups(execUser.Sgids); err != nil {
		log.Printf("E! Failed to set groups: %v", err)
		return err
	}

	if err := setGid(execUser.Gid); err != nil {
		log.Printf("E! Failed to set gid: %v", err)
		return err
	}

	if err := setUid(execUser.Uid); err != nil {
		log.Printf("E! Failed to set uid: %v", err)
		return err
	}

	if err := os.Setenv("HOME", execUser.Home); err != nil {
		log.Printf("E! Failed to set HOME: %v", err)
		return err
	}
	log.Printf("I! Set HOME: %v", execUser.Home)

	return nil
}

// getRunAsExecUser returns the new user uid and gid
func getRunAsExecUser(runasuser string) (*user.ExecUser, error) {
	currExecUser := user.ExecUser{
		Uid:  syscall.Getuid(),
		Gid:  syscall.Getgid(),
		Home: "/root",
	}
	newUser, err := user.GetExecUserPath(runasuser, &currExecUser, "/etc/passwd", "/etc/group")
	if err != nil {
		log.Printf("E! Failed to get newUser: %v", err)
		return nil, err
	}
	return newUser, nil
}

// ChangeUser allow customers to run the collector in selected user
// by default it ran as 'aoc' user but can be set by environment variable
func ChangeUser() (string, error) {

	_, err := user.LookupUser(aocUserName)
	if err != nil {
		log.Printf("E! User %s does not exist: %v", aocUserName, err)
		return "root", err
	}

	execUser, err := getRunAsExecUser(aocUserName)
	if err != nil {
		log.Printf("E! Failed to getRunAsExecUser: %v", err)
		return aocUserName, err
	}

	changeFileOwner(execUser.Uid, execUser.Gid)

	if err := switchUser(execUser); err != nil {
		log.Printf("E! failed switching to %q: %v", aocUserName, err)
		return aocUserName, err
	}

	return aocUserName, nil
}

// changeFileOwner changes the folder to new user as owner
func changeFileOwner(uid int, gid int) {
	if err := chownRecursive(uid, gid, filepath.Join(getInstallPath(), "logs")); err != nil {
		log.Printf("E! Change ownership of %slogs: %v", getInstallPath(), err)
	}

	if err := chownRecursive(uid, gid, filepath.Join(getInstallPath(), "etc")); err != nil {
		log.Printf("E! Change ownership of %setc: %v", getInstallPath(), err)
	}

	if err := chownRecursive(uid, gid, filepath.Join(getInstallPath(), "var")); err != nil {
		log.Printf("E! Change ownership of %svar: %v", getInstallPath(), err)
	}

	log.Printf("I! Change ownership to %v:%v", uid, gid)
}

func getInstallPath() string {
	return defaultInstallPath
}

func chownRecursive(uid, gid int, dir string) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmode := info.Mode()
		if fmode.IsRegular() {
			if fmode&os.ModeSetuid != 0 || fmode&os.ModeSetgid != 0 {
				return nil
			}

			if fmode.Perm()&0111 != 0 {
				return nil
			}

			if fmode.Perm()&0002 != 0 {
				return nil
			}
		}

		if fmode&os.ModeSymlink != 0 {
			return nil
		}

		if err := chown(path, uid, gid); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error change owner of dir %s to %v:%v due to error: %w", dir, uid, gid, err)
	}
	return nil
}
