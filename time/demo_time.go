package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	var temp int64
	temp = 10000
	fmt.Print(temp)
	d := time.Duration(time.Millisecond * 1000)
	newTimer := time.NewTimer(d)
	go func() {
		for {
			<-newTimer.C
			fmt.Println("开始")
			newTimer.Reset(time.Millisecond * 1000)
		}
	}()

	lis, e := net.Listen("tcp", ":8090")

	if e != nil {
		log.Fatal(fmt.Errorf("start fail:%s", e.Error()))
		return
	}
	con, e := lis.Accept()
	rpc.ServeConn(con)
}
