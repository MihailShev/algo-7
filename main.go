package main

import (
	binaryTree "algo-7/binary-tree"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	tree := binaryTree.Tree{}

	keys := genUniqRandom(1, 100)

	for _, v := range keys {
		tree.Add(v, v)
	}

	for _, v := range keys {
		value := tree.Search(v)

		if  value != v {
			log.Fatal("key", v, "is not equal", value)
		}
		fmt.Println("key", v, "value", value)
	}
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
	for i := size -1; i >= 0; i-- {
		j := getRandom(0, size)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func getRandom(min, max int) int {
	return min + rand.Intn(max - min)
}