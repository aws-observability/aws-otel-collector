package timberjack

import (
	"fmt"
	"os"
	"syscall"
)

var osChown = os.Chown // Keep for testing

var chown = func(name string, info os.FileInfo) error {
	//// Should not be opening and truncating the file. It should just perform the osChown operation on the existing file name.
	//// This is the primary cause of tests seeing 0 bytes.
	// f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	// if err != nil {
	// 	return err
	// }
	// f.Close()

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to get syscall.Stat_t from FileInfo for %s", name)
	}
	return osChown(name, int(stat.Uid), int(stat.Gid))
}
