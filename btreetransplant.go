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

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil && root.Left == node {
		root.Left = rplc
		root.Left.Parent = root
	} else if root.Right != nil && root.Right == node {
		root.Right = rplc
		root.Right.Parent = root
	} else {
		root.Left = BTreeTransplant(root.Left, node, rplc)
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}
	return root
}
