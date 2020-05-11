package function

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	max := max(1, 3)
	fmt.Println(max)
}
func TestSwap(t *testing.T) {
	s, s2 := swap("x", "y")
	fmt.Println(s, s2)

}

func TestDeclareFuncAsV(t *testing.T) {
	fmt.Println(declareFuncAsV(2.25))
}

func TestGetSequence(t *testing.T) {
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
func TestMethodTest(t *testing.T) {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积为", c1.getArea())
	fmt.Println("圆的半径为", c1.radius)
	fmt.Println("圆的半径为", c1.radius)

}
