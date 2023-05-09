package piscine

/*type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "1")
	rplc := &TreeNode{Data: "3"}
	root = BTreeTransplant(root, node, rplc)
	BTreeApplyInorder(root, fmt.Println)
}*/

func BTreeTransplant(root, nodepar, rplcde *TreeNode) *TreeNode {
	root := &TreeNode{Data: "04"}
	root.Left = &TreeNode{Data: "01"}
	root.Left.Right = &TreeNode{Data: "02"}
	root.Left.Right.Right = &TreeNode{Data: "03"}
	root.Right = &TreeNode{Data: "07"}
	root.Right.Left = &TreeNode{Data: "05"}
	root.Right.Right = &TreeNode{Data: "12"}
	root.Right.Right.Left = &TreeNode{Data: "10"}

	nodepar := root.Right
	rplcde := &TreeNode{Data: "55"}
	rplcde.Left = &TreeNode{Data: "60"}
	rplcde.Right = &TreeNode{Data: "33"}
	rplcde.Right.Left = &TreeNode{Data: "12"}
	rplcde.Right.Left.Left = &TreeNode{Data: "15"}

	root = BTreeTransplant(root, nodepar, rplcde)

}
