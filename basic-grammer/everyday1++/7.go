package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	var m error = nil
	var n interface{} = nil
	fmt.Println(x, y, z, k, p, m, n)
}
