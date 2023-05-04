package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(door *Door) bool {
	PrintStr("Door Opening...")
	door.state = "OPEN"
	return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("Is the Door opened?")
	return door.state == "OPEN"
}

func CloseDoor(door *Door) bool {
	PrintStr("Door Closing...")
	door.state = "CLOSED"
	return true
}

func IsDoorClose(door *Door) bool {
	PrintStr("Is the Door closed?")
	return door.state == "CLOSED"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
