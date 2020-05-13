package structure

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func simpleStruct() {
	//创建一个新的结构体
	fmt.Println(Book{"java 字符集", "梁自强", "java", 1})

	//key:value
	fmt.Println(Book{title: "java 字符", author: "家伙", bookId: 9})

}
func accessFieldForStruct() {
	var book1, book2 Book

	book1.bookId = 1
	book1.author = "增益"
	book1.subject = "分布式一致性"
	book1.title = "分布式原理"

	book2.bookId = 2
	book2.title = "搜索优化"
	book2.author = "lzq"
	book2.subject = "搜索的时机"

	printBook(book1)
	printBook(book2)
}

func printBook(b Book) {
	fmt.Println(b)
}

func printPtrOfBook(b *Book) {
	fmt.Println(*b)
}

func ptrForStruct() {
	var book1 Book

	book1.bookId = 1
	book1.author = "增益"
	book1.subject = "分布式一致性"
	book1.title = "分布式原理"

	var ptr *Book

	ptr = &book1

	printPtrOfBook(ptr)
}
