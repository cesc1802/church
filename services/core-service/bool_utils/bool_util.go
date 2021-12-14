package bool_utils

func BoolToBoolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}