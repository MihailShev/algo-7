package main

import (
	binaryTree "algo-7/binary-tree"
	"fmt"
	"math/rand"
	"time"
)

//var keys = []int{50, 17, 76, 9, 23, 14, 12, 19, 54, 72, 67}
//var keys = []int{50, 40, 30, 20, 10, 35, 65, 70, 25}
//var keys = []int{6, 7, 10, 8, 5, 4, 9, 3, 2, 1}
//var keys = []int{5, 1, 14, 11, 12, 18, 13, 17, 6, 2, 10, 15, 19, 20, 8, 16, 4, 7, 3, 9}
//var keys = []int{11, 13, 4, 14, 8, 10, 1, 3, 5, 12, 6, 9, 15, 7, 2}
func main() {
	tree := binaryTree.AVL{}

	keys := genUniqRandom(1, 2000)
	fmt.Println(keys)
	start := time.Now()
	for _, v := range keys {
		tree.Insert(v, v)
		if !tree.IsBalanced() {
			panic("tree is not balanced")
		}
		//fmt.Println("size", tree.Size())
		//fmt.Println(tree.String())
	}
	fmt.Println(tree.String())
	for _, v := range keys {
		fmt.Println("remove", v)
		tree.Remove(v)
		if !tree.IsBalanced() {
			panic("tree is not balanced")
		}
		fmt.Println("size", tree.Size())
		fmt.Println(tree.String())
	}
	stop := time.Since(start)

	fmt.Println("after removed")
	fmt.Println("size", tree.Size())
	fmt.Println(tree.String())
	fmt.Println("Execution time", stop)
}

func genUniqRandom(from, to int) []int {
	i := from
	res := make([]int, 0)
	for i <= to {
		res = append(res, i)
		i++
	}

	return mix(res)
}

func mix(arr []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	size := len(arr)
	for i := size - 1; i >= 0; i-- {
		j := getRandom(0, size)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func getRandom(min, max int) int {
	return min + rand.Intn(max-min)
}
