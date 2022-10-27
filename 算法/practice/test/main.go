package main

import "fmt"

type demo struct {
	Val int
}

func (d *demo) change() {
	d = nil // 对方法接收器的分配仅传播给被调用者，而不传播给调用者
	d.myVal()
}

func (d *demo) myVal() {
	fmt.Printf("my val: %#v\n", d)
}

func (d demo) change2() {
	d = demo{} // 对方法接收器的分配不会传播到其他调用
	d.myVal()
}

func (d *demo) change3() {
	d.Val = 3
	d.myVal()
}

func main() {
	d := &demo{}
	d.myVal()
	d.change()
	d.myVal()
	d.Val = 2
	d.myVal()
	d.change2()
	d.myVal()
	d.change3()
	d.myVal()
}
