package ecs

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	kernelMagicCodeNotSet      = int64(9223372036854771712) // infinity magic number for cgroup: https://unix.stackexchange.com/questions/420906/what-is-the-value-for-the-cgroups-limit-in-bytes-if-the-memory-is-not-restricte
	ecsInstanceMountConfigPath = "/proc/self/mountinfo"
)

type cgroupScanner struct {
	mountPoint string
}

func newCGroupScanner(mountConfigPath string, logger *zap.Logger) (c *cgroupScanner) {
	mp, err := getCGroupMountPoint(mountConfigPath)
	if err != nil {
		logger.Warn("Fallback to /cgroup, because it failed to get the cgroup mount point, error: ", zap.Error(err))
		mp = "/cgroup"
	}

	c = &cgroupScanner{
		mountPoint: mp,
	}
	return c
}
func newCGroupScannerForContainer(logger *zap.Logger) *cgroupScanner {
	return newCGroupScanner(ecsInstanceMountConfigPath, logger)
}

func (c *cgroupScanner) getCPUReserved(taskID string, clusterName string, logger *zap.Logger) int64 {
	cpuPath, err := getCGroupPathForTask(c.mountPoint, "cpu", taskID, clusterName)
	if err != nil {
		logger.Warn("Failed to get cpu cgroup path for task: ", zap.Error(err))
		return int64(0)
	}

	// check if hard limit is configured
	if cfsQuota, err := readInt64(cpuPath, "cpu.cfs_quota_us"); err == nil && cfsQuota != -1 {
		if cfsPeriod, err := readInt64(cpuPath, "cpu.cfs_period_us"); err == nil && cfsPeriod > 0 {
			return int64(math.Ceil(float64(1024*cfsQuota) / float64(cfsPeriod)))
		}
	}

	if shares, err := readInt64(cpuPath, "cpu.shares"); err == nil {
		return shares
	}

	return int64(0)
}

func (c *cgroupScanner) getMEMReserved(taskID string, clusterName string, containers []ECSContainer, logger *zap.Logger) int64 {
	memPath, err := getCGroupPathForTask(c.mountPoint, "memory", taskID, clusterName)
	if err != nil {
		logger.Warn("Failed to get memory cgroup path for task: %v", zap.Error(err))
		return int64(0)
	}

	if memReserved, err := readInt64(memPath, "memory.limit_in_bytes"); err == nil && memReserved != kernelMagicCodeNotSet {
		return memReserved
	}

	// sum the containers' memory if the task's memory limit is not configured
	sum := int64(0)
	for _, container := range containers {
		containerPath := path.Join(memPath, container.DockerId)

		//soft limit first

		if softLimit, err := readInt64(containerPath, "memory.soft_limit_in_bytes"); err == nil && softLimit != kernelMagicCodeNotSet {
			sum += softLimit
			continue
		}

		// try hard limit when soft limit is not configured
		if hardLimit, err := readInt64(containerPath, "memory.limit_in_bytes"); err == nil && hardLimit != kernelMagicCodeNotSet {
			sum += hardLimit
		}
	}
	return sum
}

func readString(dirpath string, file string) (string, error) {
	cgroupFile := path.Join(dirpath, file)

	// Read
	out, err := ioutil.ReadFile(cgroupFile)
	if err != nil {
		// Ignore non-existent files
		log.Printf("W! readString: Failed to read %q: %s", cgroupFile, err)
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func readInt64(dirpath string, file string) (int64, error) {
	out, err := readString(dirpath, file)
	if err != nil {
		return 0, err
	}

	if out == "" || out == "max" {
		return 0, err
	}

	val, err := strconv.ParseInt(out, 10, 64)
	if err != nil {
		log.Printf("W! readInt64: Failed to parse int %q from file %q: %s", out, path.Join(dirpath, file), err)
		return 0, err
	}

	return val, nil
}
func getCGroupMountPoint(mountConfigPath string) (string, error) {
	f, err := os.Open(mountConfigPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		var (
			text   = scanner.Text()
			fields = strings.Split(text, " ")
			// safe as mountinfo encodes mountpoints with spaces as \040.
			// an example: 26 22 0:23 / /cgroup/cpu rw,relatime - cgroup cgroup rw,cpu
			index               = strings.Index(text, " - ")
			postSeparatorFields = strings.Fields(text[index+3:])
			numPostFields       = len(postSeparatorFields)
		)
		// this is an error as we can't detect if the mount is for "cgroup"
		if numPostFields == 0 {
			return "", fmt.Errorf("Found no fields post '-' in %q", text)
		}
		if postSeparatorFields[0] == "cgroup" {
			// check that the mount is properly formated.
			if numPostFields < 3 {
				return "", fmt.Errorf("Error found less than 3 fields post '-' in %q", text)
			}
			return filepath.Dir(fields[4]), nil
		}
	}
	return "", fmt.Errorf("mount point not existed")
}

func getCGroupPathForTask(cgroupMount, controller, taskID, clusterName string) (string, error) {
	taskPath := path.Join(cgroupMount, controller, "ecs", taskID)
	if _, err := os.Stat(taskPath); os.IsNotExist(err) {
		// Task cgroup path does not exist, fallback to try legacy Task cgroup path,
		// legacy cgroup path of task with new format ARN used to contain cluster name,
		// before ECS Agent PR https://github.com/aws/amazon-ecs-agent/pull/2497/
		taskPath = path.Join(cgroupMount, controller, "ecs", clusterName, taskID)
		if _, err := os.Stat(taskPath); os.IsNotExist(err) {
			return "", fmt.Errorf("CGroup Path %q does not exist", taskPath)
		}
	}
	return taskPath, nil
}
