package main

import "fmt"

type person struct {
	Name string
	Sex  string
	Age  int
}

/*
同一种类型可以同时定义
*/
type person_1 struct {
	Name, Sex string
	Age       int
}

//嵌套结构
type person_nest struct {
	Name string
	Age  int
	// 字段是匿名对象
	Content struct {
		phone, number int
	}
}

/**
因为是匿名的字段，所以初始化的时候，必须按照顺序来，不能颠倒乱序
*/
type person_anonymous struct {
	string
	int
}

func main() {
	/*p := person{}
	p.Age = 20
	p.Name = "jack"
	fmt.Println("raw -> ", p)
	updateAge(p)
	fmt.Println("suffix -> ", p)*/

	//initPerson()

	// 传递地址
	/*	p := person{
			Age:  20,
			Name: "jack",
		}
		fmt.Println("prefix ", p)
		updatePersonByPoint(&p)
		fmt.Println("suffix ", p)*/

	// 直接在构造的时候，就取出这个对象的地址
	/*	p := &person{
			Age:  20,
			Name: "jack",
		}
		fmt.Println("prefix ", p)
		updatePersonByPoint(p)
		fmt.Println("suffix ", p)*/

	compareExample()

}

/*
比较
*/
func compareExample() {
	p := person{
		Name: "jack",
		Age:  20,
		Sex:  "gender",
	}

	p1 := person_1{
		Name: "jack",
		Age:  20,
		Sex:  "gender",
	}

	// p == p1  这两个不能直接比较，因为 person和person1虽然内容，字段相同，
	// 但是属于2种不同的类型，不同的类型不能在一起比较
	//fmt.Println(p == p1 )  // error
	fmt.Println(p)
	fmt.Println(p1)
}

func nestPersonInit() {
	//初始化
	p := person_nest{
		Name: "jack",
		Age:  20,
	}
	// 对于嵌套的匿名类型的字段，初始化只能在外面调用初始化
	p.Content.number = 100
	p.Content.phone = 1398878

	fmt.Println("raw -> ", p)
	// 引用传递
	updateNest(&p)
	fmt.Println("suffix -> ", p)

}

func updateNest(p *person_nest) {
	p.Age = 1
	p.Content.number = 119

	fmt.Println(p)
}

func anonymousStruct1() {
	// 正常直接取地址符号，作为参数，可以引用传递
	a := &struct {
		Name string
		Age  int
	}{
		Name: "jack",
		Age:  20,
	}
	fmt.Println(a)
}

func anonymousStruct() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "jack",
		Age:  20,
	}
	fmt.Println(a)
}

/**
通过指针指向的地址，传递参数，可以修改原始的值
*/
func updatePersonByPoint(p *person) {
	p.Age = 1
	fmt.Println("update ", p)
}

func initPerson() {
	p := person{
		Name: "jack",
		Age:  20,
	}

	fmt.Println(p)
}

/**
1, 此时传递的是person的值拷贝的信息，并不是原始的指针地址
在函数中更改的值，并不能更新原始地址的值
*/
func updateAge(per person) {
	per.Age = 30
	fmt.Println("updateAge", per)
	fmt.Println("%p\n", per)
}
