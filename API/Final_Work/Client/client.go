//Client Service in Go.

package main

import (
  "fmt"
  "net/rpc"
  "log"
  "strconv"
  "encoding/json"
  "os"
  "io/ioutil"
)

type Config struct {
  ServerId  string
  Protocol  string
  IpAddress  string
  Port  int
  Methods interface{}
}

func main(){
  configfile, e := ioutil.ReadFile(os.Args[1])
  if e != nil{
    fmt.Printf("Config file error : %v\n", e)
    os.Exit(1)
  }
  var configobj Config
  json.Unmarshal(configfile, &configobj)
  client_service, error := rpc.Dial(configobj.Protocol, configobj.IpAddress + ":" + strconv.Itoa(configobj.Port))
  if error != nil{
    log.Fatal("dialing:", error)
  }
  var result []byte
  for counter := 2; counter < len(os.Args); counter++{
    user_command := os.Args[counter]
    user := []byte(user_command)
    error = client_service.Call("Listings.Start", user, &result)
    if error != nil{
      log.Fatal("Connection to Server Terminated")
    }
    //Writing this just to make sure that interface is converted
    fmt.Printf(string(result))
  }
}
