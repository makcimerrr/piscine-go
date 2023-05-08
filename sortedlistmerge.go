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
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}

type NodeI struct {
	Data int
	Next *NodeI
}*/

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Fusion des deux listes
	var mergedList *NodeI
	if n1.Data < n2.Data {
		mergedList = n1
		mergedList.Next = SortedListMerge(n1.Next, n2)
	} else {
		mergedList = n2
		mergedList.Next = SortedListMerge(n1, n2.Next)
	}

	return mergedList
}
