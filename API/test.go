package main

import (
  "encoding/json"
  "fmt"
  //"os"
)

type User struct{
  Method string
  Params json.RawMessage
  Id int
}

func main(){
  user := []byte(`{"method": "lookup", "params" : ["mihir", "kelkar"], "id": 22}`)
  //obj, err := json.Marshal(user)
  //if err != nil{
  //  fmt.Println(err)
  //  return
  //}
  fmt.Println(string(user))
  var result interface{}
  error := json.Unmarshal(user, &result)
  if error != nil{
    fmt.Println("This is the error")
  }
  //fmt.Println(result)
  res, _:= result.(map[string]interface{})
  fmt.Println(res["params"])
}
