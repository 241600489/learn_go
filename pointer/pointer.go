package pointer

import "fmt"

//定义指针变量。
//为指针变量赋值。
//访问指针变量中指向地址的值。
func simplePointer() {
	var a int = 10

	var ip *int

	ip = &a

	fmt.Println("a 的地址是", &a)

	fmt.Println("ip 的值为", ip)

	fmt.Println("ip 指针指向的值", *ip)

}

//
//当一个指针被定义后没有分配到任何变量时，它的值为 nil。
//
//nil 指针也称为空指针。
//
//nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
//
//一个指针变量通常缩写为 ptr。

func nilPointer() {
	var ptr *int

	fmt.Printf("ptr 的值为%x\n", ptr)

	if ptr == nil {
		fmt.Println("ptr 是 nil")
	} else {
		fmt.Println("ptr 不是 nil")
	}

}

const MAX int = 3

func pointerOfArray() {
	a := []int{10, 100, 1000}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d]=%d\n", i, *ptr[i])
	}
}

//指针的指针
func pointerOfPointer() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	//指向a的地址
	ptr = &a
	// 指向指针 ptr 地址
	pptr = &ptr
	//获取pptr 的值
	fmt.Printf("变量a=%d\n", a)
	fmt.Printf("指针变量*ptr=%d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr=%d\n", **pptr)
}

func swap() {
	var x, y int = 9, 8

	y, x = x, y

	fmt.Printf("x=%d,y=%d\n", x, y)
}

type Square struct {
	layer int
}

func areaForSquare(c *Square) int {
	c.layer = 9
	return c.layer * c.layer
}
