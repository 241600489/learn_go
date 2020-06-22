package io

import "os"

func createAndWriteToFile() {
	create, err := os.Create("./test.log")
	if err != nil {
		println(err)
	}
	a := "nihao"

	create.Write([]byte(a))
	create.Close()
}
