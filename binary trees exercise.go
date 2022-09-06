package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		if t.Left != nil {
			Walk(t.Left, ch)
		}
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		ch <- t.Value
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	var t1vals, t2vals []int

	for i := 0; i < 10; i++ {
		v1, v2 := <-c1, <-c2
		t1vals = append(t1vals, v1)
		t2vals = append(t2vals, v2)
	}

	for i := 0; i < 10; i++ {
		matched := false
		for j := 0; j < 10; j++ {
			if t1vals[i] == t2vals[j] {				
				matched = true
			}
		}
		if(!matched){
			return false
		}
	}

	return true
}

func main() {
	t := tree.New(1)
	s := t.String()
	fmt.Print(s)

	c := make(chan int)
	go Walk(t, c)

	fmt.Println("\ntree:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
