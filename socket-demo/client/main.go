package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

//def _rpc(self, method, *args):
//    data = {'id': self.id,
//            'method': method,
//            'params': args}
//    request = json.dumps(data)
//    self.client.write(request+'\n')
//    self.client.flush()
//    response = self.client.readline()
//    self.id += 1
//    result = json.loads(response)
//    if result['error'] is not None:
//      print(result['error'])
//    # namedtuple doesn't work with unicode keys.
//    return Result(id=result['id'], result=result['result'],
//                  error=result['error'], )

type result struct {
	//'id,result,error'
	Id int `json:"id"`
	Result string `json:"result"`
	Error string `json:"error"`
}

type data struct {
	Id int `json:"id"`
	Method string `json:"method"`
	Params []string `json:"params"`
}

func main() {
	conn, err := net.Dial("tcp", "192.168.110.78:43329")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("start to test")

	ll := make([]string, 0)

	d := data{
		Id:     1,
		Method: "getU3VCameraStatus",
		Params: ll,
	}

	reqStr, err := json.Marshal(d)
	if err != nil {
		fmt.Println("json err:", err)
	}

	writer := bufio.NewWriter(conn)
	writer.WriteString(string(reqStr) + "\n")
	writer.Flush()

	//fmt.Fprintf(conn, "%s\n", reqStr)
	//line, prefix, err := bufio.NewReader(conn).ReadLine()
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	//var ret result
	//err = json.Unmarshal(res, &ret)
	//fmt.Fprintf(conn, "Jimmy!\n")
}