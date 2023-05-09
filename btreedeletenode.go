package main

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

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	BTreeApplyInorder(root, fmt.Println)
	root = BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	BTreeApplyInorder(root, fmt.Println)
}*/

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// Cas où le nœud à supprimer est une feuille
	if node.Left == nil && node.Right == nil {
		if node == root {
			// Cas où le nœud à supprimer est la racine de l'arbre
			return nil
		}
		// Cas où le nœud à supprimer est un enfant d'un autre nœud
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left != nil && node.Right == nil {
		// Cas où le nœud à supprimer a un enfant gauche
		if node == root {
			root = node.Left
			root.Parent = nil
		} else {
			if node.Parent.Left == node {
				node.Parent.Left = node.Left
			} else {
				node.Parent.Right = node.Left
			}
			node.Left.Parent = node.Parent
		}
	} else if node.Left == nil && node.Right != nil {
		// Cas où le nœud à supprimer a un enfant droit
		if node == root {
			root = node.Right
			root.Parent = nil
		} else {
			if node.Parent.Left == node {
				node.Parent.Left = node.Right
			} else {
				node.Parent.Right = node.Right
			}
			node.Right.Parent = node.Parent
		}
	} else {
		// Cas où le nœud à supprimer a deux enfants
		// On cherche le plus petit élément dans le sous-arbre droit
		// du nœud à supprimer
		successor := BTreeMin(node.Right)
		node.Data = successor.Data
		root = BTreeDeleteNode(root, successor)
	}
	return root
}
