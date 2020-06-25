package io

import (
	"bytes"
	"fmt"
)

func byteBuffer() {
	containerRaw := make([]byte, 0, 10)
	buf := bytes.NewBuffer(containerRaw)
	a := "nihaonihaonihao"
	b := []byte(a)
	writeCount, _ := buf.Write(b)

	fmt.Printf("buf's cap:%d,len:%d,have write count:%d", buf.Cap(), buf.Len(), writeCount)

}
