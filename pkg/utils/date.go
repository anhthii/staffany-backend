package utils

import (
	"strconv"
	"strings"
)

func DateStringToInt(date string) uint64 {
	removedDashStr := strings.ReplaceAll(date, "-", "")

	i, err := strconv.ParseUint(removedDashStr, 10, 64)
	if err != nil {
		return 0
	}

	return i
}
