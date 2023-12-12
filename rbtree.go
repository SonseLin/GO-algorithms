package main

import (
	"fmt"
	"math/rand"
)

type redBlackTree struct {
	value int
	color string
	left  *redBlackTree
	right *redBlackTree
}

func createNewTree(value int) redBlackTree {
	return redBlackTree{left: nil, right: nil, value: value, color: "black"}
}

func (rbtree *redBlackTree) insert(value int) {
	rbt := rbtree
	for (rbt.right != nil && rbt.value <= value) || (rbt.left != nil && rbt.value > value) {
		if rbt.value <= value {
			if rbt.color == "red" && rbt.left != nil {
				rbt.right.color = "black"
			}
			rbt = rbt.right
		} else {
			if rbt.color == "red" && rbt.right != nil {
				rbt.left.color = "black"
			}
			rbt = rbt.left
		}
	}
	newNode := createNewTree(value)
	if rbt.color == "black" {
		newNode.color = "red"
	}
	if value > rbt.value {
		rbt.right = &newNode
	} else {
		rbt.left = &newNode
	}
}

func (rbtree *redBlackTree) printTree() {
	if rbtree.right != nil {
		rbtree.right.printTree()
	}
	fmt.Println(rbtree.value, rbtree.color)
	if rbtree.left != nil {
		rbtree.left.printTree()
	}
}

func gettingValues(rbt *redBlackTree, values *[]int) {
	if rbt.right != nil {
		gettingValues(rbt.right, values)
	}
	*values = append(*values, rbt.value)
	if rbt.left != nil {
		gettingValues(rbt.left, values)
	}
}

func main() {
	rbt := createNewTree(0)
	for i := 0; i < 10; i++ {
		if i != 0 {
			rbt.insert(i)
		}
	}
	var values []int
	gettingValues(&rbt, &values)
	balancedRbt := createNewTree(values[len(values)/2])
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
	fmt.Println(values)
	for _, v := range values {
		if v != balancedRbt.value {
			balancedRbt.insert(v)
		}
	}
	fmt.Println("center is", balancedRbt.value)
	fmt.Println("balanced tree")
	balancedRbt.printTree()
}
