package varialble

import "fmt"

//函数内定义的变量称为局部变量
//函数外定义的变量称为全局变量
//函数定义中的变量称为形式参数
func localVariableDemo() {
	//声明局部变量
	var a, b, c int
	//初始化参数
	a = 10
	b = 20
	c = a + b
	fmt.Printf("结果：a=%d,b=%d and c=%d\n", a, b, c)

}

var g = 20

func globalVariableDemo() {
	g = 10
	fmt.Printf("结果：g=%d\n", g)
}

//形式参数会作为函数的局部变量来使用。

func formalParameters(x int) int {
	return x + 1

}
