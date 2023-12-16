package main

import (
	"errors"
	"fmt"
)

type hashBookValue struct {
	value []int
}

func (hbv *hashBookValue) String() string {
	return fmt.Sprintf("includes node%v", hbv.value)
}

func hashFunc(value, size int) int {
	return value % size
}

func processLoop(from, to, by int, hb map[int]*hashBookValue, size int, fun func(map[int]*hashBookValue, int, int)) {
	for i := from; i < to; i += by {
		fun(hb, size, i)
	}
}

func insertPassCollision(hashBook map[int]*hashBookValue, size, value int) {
	place := hashFunc(value, size)
	v, ok := hashBook[place]
	if ok == true {
		v.value = append(v.value, place)
		hashBook[place].value = append(hashBook[place].value, value)
	} else {
		hashBook[place] = &hashBookValue{value: []int{value}}
	}
}

func searchForKey(key int, hb map[int]*hashBookValue, size int) (hashBookValue, error) {
	if size <= key {
		return hashBookValue{value: []int{0}}, errors.New(fmt.Sprintf("key %d is beyond than size %d", key, size))
	} else {
		v, ok := hb[key]
		if ok == false {
			errorMessage := fmt.Sprintf("no key %d in hashmap", key)
			return hashBookValue{value: []int{0}}, errors.New(errorMessage)
		} else {
			return *v, nil
		}
	}
}

func deleteForKey(key int, hb map[int]*hashBookValue) {
	delete(hb, key)
}

func main() {
	size := 25
	hashBook := make(map[int]*hashBookValue, size)
	processLoop(50, 400, 13, hashBook, size, insertPassCollision)
	for i, v := range hashBook {
		fmt.Printf("Hash key [%d] %+v\n", i, v)
	}
	v, err := searchForKey(23, hashBook, size)
	if err == nil {
		fmt.Println(v.value)
	} else {
		fmt.Println(err)
	}
	deleteForKey(23, hashBook)
	v, err = searchForKey(23, hashBook, size)
	if err == nil {
		fmt.Println(v.value)
	} else {
		fmt.Println(err)
	}
}
