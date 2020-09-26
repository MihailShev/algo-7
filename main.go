package main

import (
	binaryTree "algo-7/binary-tree"
	"fmt"
	"math/rand"
	"time"
)

//var keys = []int{50, 17, 76, 9, 23, 14, 12, 19, 54, 72, 67}
var keys = []int{50, 40, 30, 20, 10, 35, 65, 70, 45, 43, 40}

func main() {
	tree := binaryTree.AVL{}

	//keys := genUniqRandom(1, 50)

	for _, v := range keys {
		tree.Insert(v, v)
		fmt.Println(tree.String())

		fmt.Printf("\n=================\n")
	}

	fmt.Println(tree.String())

	//for _, v := range keys {
	//	value := tree.Search(v)
	//
	//	if  value != v {
	//		log.Fatal("key", v, "is not equal", value)
	//	}
	//	fmt.Println("key", v, "value", value)
	//}

	//fmt.Println("size", tree.Size(), "maxDeep")
	//fmt.Println(tree.String())
	//removed := tree.Remove(50)
	//fmt.Println("removed: ", removed)
	//fmt.Println("size: ", tree.Size(), "maxDeep")
	//fmt.Println(tree.String())

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
