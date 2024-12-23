package utils

import "bytes"

func PadLeft(s string, maxLength int, fill byte) string {
	if len(s) >= maxLength {
		return s
	}

	return string(bytes.Repeat([]byte{fill}, maxLength-len(s))) + s
}
