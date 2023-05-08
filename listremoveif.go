package piscine

/*func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Tail != nil {
		l.Tail.Next = n
		l.Tail = n
	} else {
		l.Head = n
		l.Tail = n
	}
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListRemoveIf(l *List, data_ref interface{}) {
	var n *NodeL = nil
	m := l.Head
	for m != nil {
		if m.Data == data_ref {
			if n == nil {
				l.Head = m.Next
			} else {
				n.Next = m.Next
			}
			if m.Next == nil {
				l.Tail = n
			}
		} else {
			n = m
		}
		m = m.Next
	}
}
