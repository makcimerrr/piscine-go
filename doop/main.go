package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := parseInt64(os.Args[1])
	if arg1 == nil {
		return
	}

	arg2 := os.Args[2]

	arg3 := parseInt64(os.Args[3])
	if arg3 == nil {
		return
	}

	var result int64

	switch arg2 {
	case "+":
		result = *arg1 + *arg3
	case "-":
		result = *arg1 - *arg3
	case "*":
		result = *arg1 * *arg3
	case "/":
		if *arg3 == 0 {
			writeString([]byte("No division by 0\n"))
			return
		}
		result = *arg1 / *arg3
	case "%":
		if *arg3 == 0 {
			writeString([]byte("No modulo by 0\n"))
			return
		}
		result = *arg1 % *arg3
	default:
		return
	}

	if arg2 == "+" && ((*arg1 > 0 && *arg3 > 0 && result < 0) || (*arg1 < 0 && *arg3 < 0 && result > 0)) {
		return
	}
	if arg2 == "-" && ((*arg1 > 0 && *arg3 < 0 && result < 0) || (*arg1 < 0 && *arg3 > 0 && result > 0)) {
		return
	}
	if arg2 == "*" && ((*arg1 > 0 && *arg3 > 0 && result > 0) || (*arg1 < 0 && *arg3 < 0 && result > 0)) {
		return
	}
	if arg2 == "/" && ((*arg1 > 0 && *arg3 < 0 && result < 0) || (*arg1 < 0 && *arg3 > 0 && result > 0)) {
		return
	}

	writeInt64(result)
	writeByte('\n')
}

func parseInt64(s string) *int64 {
	var x int64
	var neg bool
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return nil
		}
		x = x*10 + int64(s[i]-'0')
	}
	if neg {
		x = -x
	}
	return &x
}

func writeString(s []byte) {
	os.Stdout.Write(s)
}

func writeByte(b byte) {
	os.Stdout.Write([]byte{b})
}

func writeInt64(n int64) {
	if n < 0 {
		writeByte('-')
		n = -n
	}
	if n > 9 {
		writeInt64(n / 10)
	}
	writeByte(byte(n%10) + '0')
}
