package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
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
}*/

func ListReverse(l *List) {
	var n *NodeL
	head := l.Head
	l.Tail = l.Head

	for head != nil {
		next := head.Next
		head.Next = n
		n = head
		head = next
	}

	l.Head = n
}
