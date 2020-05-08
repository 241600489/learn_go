package demo

import "fmt"

func dist() {
	fmt.Println("nihao")
	var a int16 = 7
	fmt.Println(a)
	var b = "nihaog"
	fmt.Println(b)

}

func va() {
	var a string = "niaho"
	fmt.Println(a)
	var b = "wohaini"
	fmt.Println(b)
	c := "niaho"
	fmt.Println(c)
}

const (
	a = iota
	b
	c
	d = "xiaozeng"
	e
	f = "xiaoliang"
	g
)
const (
	i = 1 << iota
	j = 3 << iota
	k //k=3<<2
	l //l=3<<3
)

func constDemo() {
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

func ifDemo() {
	a := 1
	var b = 5
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<b")
	}
}
func findPrimeNumber() {
	var count int
	var primeNumberConstainer = make([]int, 1)

	for count < 100 {
		var flag bool
		count++
		for tem := 2; tem < count; tem++ {
			if count%tem != 0 {
				flag = true
			}
		}
		if flag {
			primeNumberConstainer = append(primeNumberConstainer, count)
		}
	}
	fmt.Println(primeNumberConstainer)
}
func switchDemo() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T", i)
	case int:
		fmt.Println("x 是整型")

	default:
		fmt.Printf("x 啥也不是")
	}
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	case 60:
		grade = "D"
	default:
		grade = "E"
	}
	switch {
	case grade == "B":
		fmt.Println("成绩良好")
	case grade == "C":
		fmt.Println("成绩有待提高")
	case grade == "A":
		fmt.Println("成绩非常优秀，再接再厉")
	default:
		fmt.Println("革命尚未成功，同志仍需努力")
	}

}

func selectDemo() {
	var c1, c2, c3 chan int32
	var i1, i2 int32
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, "from c1")
	case i2 = <-c2:
		fmt.Println("received ", i2, "from c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received ", i3, "from c3")
		} else {
			fmt.Printf("c3 is closed")
		}
	default:
		fmt.Printf("no communication \n")

	}
}
