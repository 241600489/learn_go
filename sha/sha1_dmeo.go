package sha

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

type ID uint64

func sha1Demo() {
	var ips = [3]string{"127.0.0.1:7980", "127.0.0.1:7981", "127.0.0.1:7982"}
	for _, s := range ips {
		var b []byte
		b = append(b, []byte(s)...)
		hash := sha1.Sum(b)
		id := ID(binary.BigEndian.Uint64(hash[:8]))
		fmt.Println(id)
	}

}
