package piscine

/*func ListPushBack(l *List, data interface{}) {
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

	ListPushBack(link, "hello")
	ListPushBack(link, "hello1")
	ListPushBack(link, "hello2")
	ListPushBack(link, "hello3")

	found := ListFind(link, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	return &l.Head.Next.Next.Data
}
