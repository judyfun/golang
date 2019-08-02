package main

import (
	"fmt"
	"sort"
)

func main() {
	homeWork()
}

/**

 */
func homeWork() {
	a := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	b := make(map[string]int)

	s := make([]int, len(a))
	i := 0
	for k := range a {
		s[i] = k
		i++
	}

	sort.Ints(s)

	for i, v := range a {
		b[v] = i
	}

	fmt.Println(s)

	fmt.Println(a)
	fmt.Println(b)
}

/**
迭代返回的参数，i,v
i 是索引值
v 是value，v 是一个value的拷贝，改变value并不能改变原有地址里的value
*/
func entryMap() {
	sm := make([]map[int]string, 5)
	for i, v := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(i)
		fmt.Println(sm[i][1])
		fmt.Println(v)
	}
}

//迭代
func entry() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i)
		fmt.Println(v)
	}
}

//复杂的嵌套map
func complcateMap() {
	m := make(map[int]map[int]string)
	a, ok := m[1][1]

	// 需要判断下，嵌套的是否被初始化了，默认只初始化第一层的数据
	if !ok {
		m[1] = make(map[int]string)
	}
	m[1][1] = "hello"

	fmt.Println(m)
	fmt.Println(a)

}

func deleteMap() {
	m := make(map[int]string)
	m[1] = "hello"
	m[2] = "world"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

}
func initMap1() {
	m := make(map[int]string)

	m[1] = "hello"
	m[1] = "hello hello"
	a := m[1]
	fmt.Println(m)
	fmt.Println(a)
}
func initMap() {
	// 声明
	var m map[int]string
	m = map[int]string{}

	m[1] = "hello"
	fmt.Println(m)
}
