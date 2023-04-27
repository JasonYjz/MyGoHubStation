package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Result struct {
	Id string `json:"id"`
	Result string `json:"result"`
	Error string `json:"error"`
}

type RpcReq struct {
	Id int `json:"id"`
	//Method string `json:"method"`
	Params []string `json:"params"`
}

func callRpcWithSynchronous() {
	client, err := net.DialTimeout("tcp", "192.168.110.78:43329", 1000*1000*1000*10)
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}
	defer client.Close()

	// 建立RPC通道
	rpcClient := jsonrpc.NewClient(client)

	ll := make([]string, 0)
	// 测试1
	// 服务器返回对象
	var response Result
	req := RpcReq{
		Id:     1,
		//Method: "getU3VCameraStatus",
		Params: ll,
	}
	//request1 := 1
	log.Println("client\t-", "Call getU3VCameraStatus method")
	// 请求数据，rpcObj对象会被填充
	rpcClient.Call("getU3VCameraStatus", req, &response) // 此处必须传入指针
	log.Println("client\t-", "Receive remote return:", response)
}

func main() {
	callRpcWithSynchronous()
}
