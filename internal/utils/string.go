package utils

const ErrorFormat = "%s: %w"

func GetStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
