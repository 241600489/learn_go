package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

func main() {
	//内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	//new返回指针。
	say := new(Say)
	rpc.RegisterName("Say1", say)
	//ResolveTCPAddr返回TCP端点的地址。
	//网络必须是TCP网络名称。
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	log.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		// todo   需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

type Say struct {
}

type VoteRequest struct {
	R int
}
type VoteResponse struct {
	VoteGranted bool
}

func (s *Say) Say(vr *VoteRequest, resp *VoteResponse) (err error) {
	log.Println("接收到请求")
	resp.VoteGranted = true
	return nil

}
