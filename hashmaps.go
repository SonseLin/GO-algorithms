package main

import "fmt"

type hashBookValue struct {
	value []int
}

func processLoop(from, to, by int, hb map[int]*hashBookValue, size int, fun func(map[int]*hashBookValue, int, int)) {
	for i := from; i < to; i += by {
		fun(hb, size, i)
	}
}

func insertPassCollision(hashBook map[int]*hashBookValue, size, value int) {
	v, ok := hashBook[value%size]
	if ok == true {
		v.value = append(v.value, value%size)
		hashBook[value%size].value = append(hashBook[value%size].value, value)
	} else {
		hashBook[value%size] = &hashBookValue{value: []int{value}}
	}
}

func main() {
	size := 25
	hashBook := make(map[int]*hashBookValue, size)
	processLoop(50, 400, 13, hashBook, size, insertPassCollision)
	for i, v := range hashBook {
		fmt.Printf("Hash key [%d] has value %v\n", i, v)
	}
}
