package conv

import "strconv"

// Conv converts a string to an int64 and returns an error if conversion fails.
func Conv(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}