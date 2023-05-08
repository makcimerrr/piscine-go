package piscine

func ListSize(l *List) int {
	size := 0
	for n := l.Head; n != nil; n = n.Next {
		size++
	}
	return size
}

/*func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}*/
