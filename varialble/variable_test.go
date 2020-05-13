package varialble

import (
	"fmt"
	"testing"
)

func TestLocalVariableDemo(t *testing.T) {
	localVariableDemo()
}

func TestGlobalVariableDemo(t *testing.T) {
	globalVariableDemo()
}

func TestFormalParameters(t *testing.T) {
	result := formalParameters(9)
	fmt.Printf("result:%d", result)
}
