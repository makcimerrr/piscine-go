package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := make([]byte, 0, len(args)*20)
	for _, arg := range args {
		result = append(result, []byte(arg)...)
		result = append(result, '\n')
	}
	return string(result)
}
