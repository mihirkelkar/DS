//Client Service in Go.

package main

import (
  "fmt"
  "net/rpc"
  "log"
)

func main(){
  user := `{"method": "insert", "params" : ["m", "k", {"a" : 1, "b" : 1}], "id": 22}`
  user_one :=  `{"method": "listIDs", "params" : ["m", "k"], "id": 23}`
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
  //data, _ := result.(string)
  fmt.Printf(string(result))
  error = client_service.Call("Listings.Start", user_one, &result)
  if error != nil{
    log.Fatal("Connection to server Terminated")
  }
  //Writing this just to make sure that interface is converted
  //data, _ := result.(string)
  fmt.Printf(string(result))
}
