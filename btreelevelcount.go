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
	fmt.Println(BTreeLevelCount(root))
}*/

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := BTreeLevelCount(root.Right)
	left := BTreeLevelCount(root.Left)
	if left < right {
		right = right + 1
		return right
	}
	left = left + 1
	return left
}
