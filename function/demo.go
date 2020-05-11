package function

import "math"

func max(num1, num2 int32) int32 {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
func swap(x, y string) (string, string) {
	return y, x
}
func declareFuncAsV(x float64) float64 {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	return getSquareRoot(x)
}

/**
 * Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
 *  匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
 *  以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下
 */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Circle struct {
	radius float64
}

func (c *Circle) getArea() float64 {
	c.radius = 5

	return 3.14 * c.radius * c.radius
}
