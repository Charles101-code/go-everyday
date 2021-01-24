package main

import "fmt"

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	//list := new([]int)
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	fmt.Println(size, max_size)
}
