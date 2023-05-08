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

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}*/

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	// Initialisation de la liste triée avec le premier élément
	n := &NodeI{Data: l.Data}
	l = l.Next
	for l != nil {
		// Recherche de la position d'insertion
		m := &NodeI{}
		z := n
		for z != nil && z.Data < l.Data {
			m = z
			z = z.Next
		}
		// Insertion du nœud dans la liste triée
		node := &NodeI{Data: l.Data, Next: z}
		if z == n {
			n = node
		} else {
			m.Next = node
		}
		l = l.Next
	}
	return n
}
