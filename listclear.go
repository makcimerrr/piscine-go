package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListClear(l *List) {
	l.Head = nil
}

/*func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	ListPushBack(link, "I")
	ListPushBack(link, 1)
	ListPushBack(link, "something")
	ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}*/
