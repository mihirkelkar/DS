//Server.go
//Mihir Kelkar
package main

import (
    "fmt"
    "net"
    "net/rpc"
    "log"
  )
//This is the part that needs to replaced by the map data_structure or by the 
//database
type Args struct{
  X, Y int
}

//Define a "class" and its "methods"

type Calculator struct{}

func (t *Calculator) Add(args *Args, reply *int) error{
    *reply = args.X + args.Y
    fmt.Println(*reply)
    return nil
  }

func main(){
  //Instnatiate a new "object"
  cal := new(Calculator)
  //Regsiter the new object
  rpc.Register(cal)
  listener, error := net.Listen("tcp" , ":1234")
  if error != nil{
    log.Fatal("Error occured while listening" + error.Error())
  }
  for{
    connection, err := listener.Accept()
    if err != nil{
      log.Fatal("Error while accepting listener request")
    } else {
      log.Println("Connection to the client service established") 
    }
    go rpc.ServeConn(connection)
    fmt.Println("Request served")
  }
}
