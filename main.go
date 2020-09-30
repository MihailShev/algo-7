package main

import (
	binaryTree "algo-7/binary-tree"
	"algo-7/utils"
	"fmt"
	"time"
)

const dataSize = 80_000

func main() {
	orderDataSet := utils.FillArray(1, dataSize)
	randomDataSet := utils.FillArrayUniqRandom(1, dataSize)

	valuesToSearchFromOrder := utils.GetRandomValueList(orderDataSet, len(orderDataSet)/10)
	valuesToRemoveFromOrder := utils.GetRandomValueList(orderDataSet, len(orderDataSet)/10)

	valuesToSearchFromRandom := utils.GetRandomValueList(randomDataSet, len(orderDataSet)/10)
	valuesToRemoveFromRandom := utils.GetRandomValueList(randomDataSet, len(orderDataSet)/10)

	tree := binaryTree.Tree{}
	avl := binaryTree.AVL{}

	fmt.Printf("\n*** Test binary tree ***\n")
	fmt.Printf("\n=== Random dataset ===\n")
	fmt.Printf("\nInsert %d elements \n", dataSize)
	test(func() {
		for _, v := range randomDataSet {
			tree.Insert(v, v)
		}
	})

	fmt.Printf("Search %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToSearchFromRandom {
			tree.Search(v)
		}
	})

	fmt.Printf("Remove %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToRemoveFromRandom {
			tree.Remove(v)
		}
	})

	tree.Clear()

	fmt.Printf("=== Order dataset ===\n")
	fmt.Printf("\nInsert %d elements \n", dataSize)
	test(func() {
		for _, v := range orderDataSet {
			tree.Insert(v, v)
		}
	})

	fmt.Printf("Search %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToSearchFromOrder {
			tree.Search(v)
		}
	})

	fmt.Printf("Remove %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToRemoveFromOrder {
			tree.Remove(v)
		}
	})

	fmt.Printf("\n*** Test AVL tree ***\n")
	fmt.Printf("\n=== Random dataset ===\n")
	fmt.Printf("\nInsert %d elements \n", dataSize)
	test(func() {
		for _, v := range randomDataSet {
			avl.Insert(v, v)
		}
	})

	fmt.Printf("Search %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToSearchFromRandom {
			avl.Search(v)
		}
	})

	fmt.Printf("Remove %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToRemoveFromRandom {
			avl.Search(v)
		}
	})

	avl.Clear()

	fmt.Printf("=== Order dataset ===\n")
	fmt.Printf("\nInsert %d elements \n", dataSize)
	test(func() {
		for _, v := range orderDataSet {
			avl.Insert(v, v)
		}
	})

	fmt.Printf("Search %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToSearchFromOrder {
			avl.Search(v)
		}
	})

	fmt.Printf("Remove %d elements \n", dataSize/10)
	test(func() {
		for _, v := range valuesToRemoveFromOrder {
			avl.Search(v)
		}
	})

	fmt.Println("Finish, press any key")
	_, _ = fmt.Scanf(" ")

}

func test(run func()) {
	start := time.Now()
	run()
	stop := time.Since(start)
	fmt.Printf("Execution time: %s\n\n", stop)
}
