package pointer

import (
	"fmt"
	"testing"
)

func TestSimplePointer(t *testing.T) {

	simplePointer()

}

func TestNilPointer(t *testing.T) {
	nilPointer()
}

func TestPointerOfArray(t *testing.T) {
	pointerOfArray()
}

func TestPointerOfPointer(t *testing.T) {
	pointerOfPointer()
}

func TestSwap(t *testing.T) {
	swap()
}

func TestSquare(t *testing.T) {
	var s Square
	s.layer = 10

	fmt.Println(areaForSquare(&s))

}
