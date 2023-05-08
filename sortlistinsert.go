package piscine

/*func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}

type NodeI struct {
	Data int
	Next *NodeI
}*/

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	curr := l
	for curr.Next != nil && curr.Next.Data < data_ref {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode

	return l
}
