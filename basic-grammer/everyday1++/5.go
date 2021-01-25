package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "steven"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "steven"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 26, m: map[string]string{"name": "charles"}}

	sm2 := struct {
		age int
		m   map[string]string "
	}{age: 26, m: map[string]string{"name": "charles"}}

	// invalid operation: sm1 == sm2
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}

/**
 * 1. 结构体只能比较是否相等，但不能比较大小
 * 2. 相同类型的结构体才能比较，结构体是否相同不但与属性类型有关，还与属性顺序有关
 * 3. 如果struct的所有成员都可以比较，则该struct就可以通过== 或 != 进行比较是否相等，比较时逐个项进行比较。
 * 4. 切片、map、函数等是不能进行比较的
 */
