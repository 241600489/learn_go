package builtn

import (
	"encoding/binary"
	"fmt"
	crc322 "hash/crc32"
)

func crc32() {
	a := "123"
	b := []byte(a)
	checkSumB := crc322.ChecksumIEEE(b)
	ab := "123"
	ac := []byte(ab)
	checkAc := crc322.ChecksumIEEE(ac)
	fmt.Println(checkSumB, checkAc)
	v := make([]byte, 10)
	binary.BigEndian.PutUint32(v, checkAc)
	fmt.Println(v)
}
