package io

import "encoding/binary"

//字节序
func endian() {
	uint64ByteCB := make([]byte, 8)
	binary.BigEndian.PutUint64(uint64ByteCB, 90)
	for _, item := range uint64ByteCB {
		print(item, ",")
	}
	println()
	uint64ByteCL := make([]byte, 8)
	binary.LittleEndian.PutUint64(uint64ByteCL, 90)
	for _, item := range uint64ByteCL {
		print(item, ",")
	}

}
