package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println(s1)

	r, err := funcMui(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func funcMui(x, y int) (sum int, err error) {
	return x + y, err
}

// make and new are difference

/**
 *	new(T) 和 make(T, args) 是 Go 语言内置函数，用来分配内存，但适用的类型不同。
 *  new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话来说，返回一个指针，
 *		该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
 *  make(T, args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的T的
 * 		引用。make() 只适用于 slice, map, channel.
 */
