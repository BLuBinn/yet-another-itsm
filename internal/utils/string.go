package utils

func GetStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
