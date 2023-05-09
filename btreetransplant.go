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
	if nodepar == nil {
		return rplcde
	}
	if nodepar.Left == rplcde {
		return root
	}
	if nodepar.Parent != nil {
		if nodepar.Parent.Left == nodepar {
			nodepar.Parent.Left = rplcde
		} else {
			nodepar.Parent.Right = rplcde
		}
	}
	if rplcde != nil {
		if nodepar.Left != nil {
			rplcde.Left = nodepar.Left
			nodepar.Left.Parent = rplcde
		}
		if nodepar.Right != nil {
			rplcde.Right = nodepar.Right
			nodepar.Right.Parent = rplcde
		}
		rplcde.Parent = nodepar.Parent
	}
	if nodepar == root {
		return rplcde
	}
	if nodepar.Parent.Left == nodepar {
		nodepar.Parent.Left = nil
	} else {
		nodepar.Parent.Right = nil
	}
	nodepar.Parent = nil
	return root
}
