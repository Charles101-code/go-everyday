package main

import "fmt"

func main() {
	slice := []int{8, 2, 4, 5}
	m := make(map[int]*int)

	// for range in loop will be create every element transcript，not point
	for key, value := range slice {
		val := value
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
