package main

import (
	"log"
	"net/rpc"
)

type VoteRequest struct {
	R int
}
type VoteResponse struct {
	VoteGranted bool
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("链接远程失败", err)
	}
	vr := VoteRequest{}
	vr.R = 0
	var resp VoteResponse
	err = client.Call("Say1.Say", vr, &resp)
	if err != nil {
		log.Fatal("调用say 失败", err)

	}
	log.Println(resp)

}
