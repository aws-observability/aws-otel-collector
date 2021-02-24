// +build linux

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
	"log"
	"os"
	"os/exec"
	gouser "os/user"
	"strconv"
	"syscall"

	"github.com/opencontainers/runc/libcontainer/user"
	"golang.org/x/sys/unix"
)

var defaultUser = "aoc"
var defaultInstallPath = "/opt/aws/aws-otel-collector/"

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
	runAsUser := getCustomUser()
	log.Printf("I! Detected runAsUser: %v", runAsUser)

	_, err := user.LookupUser(runAsUser)
	if err != nil {
		log.Printf("E! User %s does not exist: %v", runAsUser, err)
		return "root", err
	}

	if runAsUser == "root" {
		return "root", nil
	}

	execUser, err := getRunAsExecUser(runAsUser)
	if err != nil {
		log.Printf("E! Failed to getRunAsExecUser: %v", err)
		return runAsUser, err
	}

	changeFileOwner(runAsUser, execUser.Gid)

	if err := switchUser(execUser); err != nil {
		log.Printf("E! failed switching to %q: %v", runAsUser, err)
		return runAsUser, err
	}

	return runAsUser, nil
}

// changeFileOwner changes the folder to new user as owner
func changeFileOwner(runAsUser string, groupId int) {
	group, err := gouser.LookupGroupId(strconv.Itoa(groupId))
	owner := runAsUser
	if err == nil {
		owner = owner + ":" + group.Name
	} else {
		log.Printf("I! Failed to get the group name: %v, it will just change the user, but not group.", err)
	}
	log.Printf("I! Change ownership to %v", owner)

	chowncmd := exec.Command("chown", "-R", "-L", owner, getInstallPath()+"logs")
	stdoutStderr, err := chowncmd.CombinedOutput()
	if err != nil {
		log.Printf("E! Change ownership of %slogs: %s %v", getInstallPath(), stdoutStderr, err)
	}

	chowncmd = exec.Command("chown", "-R", "-L", owner, getInstallPath()+"etc")
	stdoutStderr, err = chowncmd.CombinedOutput()
	if err != nil {
		log.Printf("E! Change ownership of %setc: %s %v", getInstallPath(), stdoutStderr, err)
	}

	chowncmd = exec.Command("chown", "-R", "-L", owner, getInstallPath()+"var")
	stdoutStderr, err = chowncmd.CombinedOutput()
	if err != nil {
		log.Printf("E! Change ownership of %svar: %s %v", getInstallPath(), stdoutStderr, err)
	}
}

// getCustomUser allows to override the default user
func getCustomUser() string {
	if user, ok := os.LookupEnv("AOT_RUN_USER"); ok {
		defaultUser = user
	}
	return defaultUser
}

func getInstallPath() string {
	return defaultInstallPath
}
