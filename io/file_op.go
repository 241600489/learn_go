package io

import (
	"log"
	"os"
)

func createAndWriteToFile() {
	create, err := os.Create("./test.log")
	if err != nil {
		println(err)
	}
	a := "nihao"

	create.Write([]byte(a))
	create.Close()
}
func fileMode() {
	dest, err := os.OpenFile("./test1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("create file error", err)
	}
	a := "wohai"
	_, err1 := dest.Write([]byte(a))
	if err1 != nil {
		log.Println("write file err", err1)
	}

	dest.Close()
}
