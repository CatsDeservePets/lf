//go:build illumos || netbsd || solaris

package main

import (
	"golang.org/x/sys/unix"
)

func diskFree(wd string) string {
	var stat unix.Statvfs_t

	if err := unix.Statvfs(wd, &stat); err != nil {
		warnf("getting free disk space: %s", err)
		return ""
	}

	// Available blocks * size per block = available space in bytes
	return "df: " + humanize(int64(stat.Bavail)*int64(stat.Bsize))
}
