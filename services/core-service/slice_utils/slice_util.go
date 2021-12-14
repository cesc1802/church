package slice_utils

import "strings"

func SliceStringToString(sliceVal []string, sep string) string {
	if sep == "" {
		sep = ","
	}
	return strings.Join(sliceVal, sep)
}

func BoolToBoolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
