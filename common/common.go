package common

import (
	"errors"
	"os"
)

type CompressStats struct {
	Size int64
}

type CompressionResult struct {
	BeforeStats CompressStats
	AfterStats  CompressStats
}

func CalculateFileStats(file string) *CompressStats {
	fi, err := os.Stat(file)
	if err != nil {
		return nil
	}

	return &CompressStats{
		Size: fi.Size(),
	}
}

func FileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	return true, nil
}