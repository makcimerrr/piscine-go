package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	step := 0
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			step = step + 1
		} else {
			start = start*3 + 1
			step = step + 1
		}
	}
	return step
}

/*func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}*/
