package pkg

import (
	"golang.org/x/sys/unix"
	"os"
)

// DiskSpace returns the available space in bytes or 0 in case
// of error.
func DiskSpace() (uint64, error) {
	var stat unix.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	if err = unix.Statfs(wd, &stat); err != nil {
		return 0, err
	}
	// Available blocks * size per block = available space in bytes
	return stat.Bavail * uint64(stat.Bsize), nil
}

/*
	space, errDisk := DiskSpace()
	switch errDisk {
	case nil:
		if space == 0 {
			logger.Error().Str("address", u.AccountName).Str("stdout", out).Str("stderr", err).Msg("ran out of disk space")
		}
	default:
		logger.Error().Err(errDisk).Msg("DiskSpace")
	}
*/