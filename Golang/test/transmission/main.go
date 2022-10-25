package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//值传递
func main1() {
	jack := Person{Name: "Jack", Age: 20}
	fmt.Printf("原始变量 Person：%+v, 内存地址：%p\n", jack, &jack)
	Rename1(jack)
	fmt.Printf("调用改名函数后的原始 Person：%+v, 内存地址：%p\n", jack, &jack)
}

func Rename1(man Person) {
	man.Name = "Tom"
	fmt.Printf("改名后的 Person：%+v, 内存地址：%p\n", man, &man)
}

//引用传递
func main() {
	jack := &Person{Name: "Jack", Age: 20}
	fmt.Printf("原始变量 Person：%+v, 内存地址：%p, 指针地址：%p\n", jack, jack, &jack)
	Rename(jack)
	fmt.Printf("调用改名函数后的原始 Person：%+v, 内存地址：%p, 指针地址：%p\n", jack, jack, &jack)
}

func Rename(man *Person) {
	fmt.Printf("传到函数里面的变量 Person：%+v, 内存地址：%p, 指针地址：%p\n", man, man, &man)
	man.Name = "Tom"
	fmt.Printf("改名后的 Person：%+v, 内存地址： %p, 指针地址：%p\n", man, man, &man)
}
