//Client Service in Go.

package main

import (
  "fmt"
  "net/rpc"
  "log"
)

func main(){
  user := `{"method": "shutdown", "params" : ["tp", "kelkar", {"a" : 1, "b": 3}], "id": 22}`
  client_service, error := rpc.Dial("tcp", "127.0.0.1:1234")
  if error != nil{
    log.Fatal("dialing:", error)
  }
  var result []byte
  error = client_service.Call("Listings.Start", user, &result)
  if error != nil{
    log.Fatal("Connection to Server Terminated")
  }
  //Writing this just to make sure that interface is converted
  fmt.Printf(string(result))
}
