package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data, Next: l.Head}

	l.Head = n

	if l.Tail != nil {
		return
	} else {
		l.Tail = n
	}
}

/*func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}*/
