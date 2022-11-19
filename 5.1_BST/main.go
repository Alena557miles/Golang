package main

import "fmt"

// BST - search, insert, delete

type bstNode struct {
	value int
	left  *bstNode
	right *bstNode
}

type bst struct {
	root *bstNode
}

func (b *bst) Insert(value int) {
	b.InsertRec(b.root, value)
}

func (b *bst) InsertRec(node *bstNode, value int) *bstNode {
	if b.root == nil {
		b.root = &bstNode{value: value}
		return b.root
	}
	if node == nil {
		return &bstNode{value: value}
	}
	if value <= node.value {
		node.left = b.InsertRec(node.left, value)
	}
	if value > node.value {
		node.right = b.InsertRec(node.right, value)
	}
	return node
}

func minValued(b *bstNode) *bstNode {
	temp := b
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}
func (b *bst) Delete(value int) {
	b.DeleteNode(b.root, value)
}
func (b *bst) DeleteNode(node *bstNode, value int) *bstNode {
	if node == nil {
		return b.root
	}
	if value < node.value {
		node.left = b.DeleteNode(node.left, value)
	} else if value > node.value {
		node.right = b.DeleteNode(node.right, value)
	} else if value == node.value {
		if node.left == nil && node.right == nil {
			node = nil
			return node
		}
		if node.left == nil && node.right != nil {
			temp := node.right
			node = nil
			node = temp
			return node
		}
		if node.left != nil && node.right == nil {
			temp := node.left
			node = nil
			node = temp
			return node
		}
		tempNode := minValued(node.right)
		node.value = tempNode.value
		node.right = b.DeleteNode(node.right, tempNode.value)
	}
	return b.root
}

func (b *bst) Search(value int) error {
	node := b.SearchNode(b.root, value)
	if node == nil {
		return fmt.Errorf("Value: %d is not exist in a tree", value)
	}
	return nil
}

func (b *bst) SearchNode(node *bstNode, value int) *bstNode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.SearchNode(node.left, value)
	}
	return b.SearchNode(node.right, value)
}

func (b *bst) inorder() {
	b.inorderRec(b.root)
	fmt.Println("b.root.value", b.root.value)
}

func (b *bst) inorderRec(node *bstNode) {
	if node != nil {
		b.inorderRec(node.left)
		fmt.Println(node.value)
		b.inorderRec(node.right)
	}
}

func (b *bst) reset() {
	b.root = nil
}

func main() {
	bst := &bst{}
	e := []int{8, 3, 70}
	for _, val := range e {
		bst.Insert(val)
	}
	fmt.Printf("Printing Inorder:\n")
	bst.inorder()
	bst.Delete(70)
	// bst.reset()
	bst.inorder()

	// eg := []int{4, 5, 7, 6, -1, 99, 5}
	// for _, val := range eg {
	// 	bst.Insert(val)
	// }
	// fmt.Printf("\nPrinting Inorder:\n")
	// bst.inorder()
	// fmt.Printf("\nFinding Values:\n")
	// err := bst.Search(7)
	// if err != nil {
	// 	fmt.Printf("Value %d Not Found\n", 7)
	// } else {
	// 	fmt.Printf("Value %d Found\n", 7)
	// }
	// err = bst.Search(6)
	// if err != nil {
	// 	fmt.Printf("Value %d Not Found\n", 6)
	// } else {
	// 	fmt.Printf("Value %d Found\n", 6)
	// }
	// err = bst.Search(5)
	// if err != nil {
	// 	fmt.Printf("Value %d Not Found\n", 5)
	// } else {
	// 	fmt.Printf("Value %d Found\n", 5)
	// }
	// err = bst.Search(1)
	// if err != nil {
	// 	fmt.Printf("Value %d Not Found\n", 1)
	// } else {
	// 	fmt.Printf("Value %d Found\n", 1)
	// }
}
