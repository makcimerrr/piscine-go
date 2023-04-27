package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := make([]byte, 0, len(args)*20)
	for c, arg := range args {
		result = append(result, []byte(arg)...)
		if c < len(args)-1 {
			result = append(result, '\n')
		}
	}
	return string(result)
}
