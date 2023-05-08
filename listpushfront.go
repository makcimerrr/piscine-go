package piscine

type NodeLs struct {
	Data interface{}
	Next *NodeLs
}

type Lists struct {
	Head *NodeLs
	Tail *NodeLs
}

func ListPushFront(l *Lists, data interface{}) {
	n := &NodeLs{Data: data, Next: l.Head}

	l.Head = n

	if l.Tail != nil {
		return
	} else {
		l.Tail = n
	}
}

/*func main() {
	link := &Lists{}

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
