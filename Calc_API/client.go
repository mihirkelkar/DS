//Client Service in Go. 

package main

import (
  "fmt"
  "net/rpc"
  "log"
)

type Args struct{
  X, Y int
}

func main(){
  client_service, error := rpc.Dial("tcp", "127.0.0.1:1234")
  if error != nil{
    log.Fatal("dialing:", error)
  }

  args := &Args{7, 8}
  var reply int
  error = client_service.Call("Calculator.Add", args, &reply)
  if error != nil{
    log.Fatal("Failure from the function", error)
  }
  fmt.Printf("Result : %d+%d=%d", args.X, args.Y, reply)
}
