package main

import (
    "fmt"
    "net/rpc"
    "net"
    "log"
    "reflect"
    "encoding/json"
//    "sync/atomic"
//    "runtime"
)


//The global map, this map will evenutally be read out the DICT3 
var global_map map[string]interface{}


//The struct is registered in the main function
type Listings struct{}


func Lookup(key string, rel string) (interface{}){
  //This function is currently working
  fmt.Println("Lookup function executed")
  return global_map[key + rel]
}

func(t *Listings) Start( data string, reply *[]byte) error{
  data_byte := []byte(data)
  //fmt.Println(string(data_byte))

  //Creating a holder than can only contain the request recieved
  var function_holder interface{}
  error := json.Unmarshal(data_byte, &function_holder)
  if error != nil{
    fmt.Println("Some problem with request recieved from the client")
  }


  //Converting function holder interface to a map
  function_res, _ := function_holder.(map[string]interface{})
  switch function_res["method"]{
    case "lookup":
      //Things done here are for the sole purpose of converting 
      //params to lookup params type. 
      fmt.Println(reflect.TypeOf(function_res["params"]))
      params, _ := function_res["params"].([]interface{})
      key, relation := params[0].(string), params[1].(string)
      result := Lookup(key, relation)
      var result_map map[string]interface{}
      result_map = make(map[string]interface{})
      result_map["result"] = result
      result_map["id"] = function_res["id"]
      result_map["error"] = nil
      *reply, _ = json.Marshal(result_map)

    case "insert":
      params, _ := function_res["params"].([]interface{})
      key, relation := params[0].(string), params[1].(string)
      value := params[2] //This is a map of string to interface
      global_map[key + relation] = value
      var result_map map[string]interface{}
      result_map = make(map[string]interface{})
      result_map["result"] = "True"
      result_map["id"] = function_res["id"]
      result_map["error"] = nil
      *reply, _ = json.Marshal(result_map)

    case "insertOrUpdate":
      params, _ := function_res["params"].([]interface{})
      key, relation := params[0].(string), params[1].(string)
      value := params[2] //This is a map of string to interface
      global_map[key + relation] = value
      *reply = nil
    
    case "delete":
      params, _ := function_res["params"].([]interface{})
      key, relation := params[0].(string), params[1].(string)
      delete(global_map, key + relation)
      *reply = []byte("This has been deleted.")

    case "listKeys":
      keys := make([]string, 0, len(global_map))
      for counter := range global_map{
        keys = append(keys, counter)
      } 
      result := interface{}(keys)
      var result_map map[string]interface{}
      result_map = make(map[string]interface{})
      result_map["result"] = result
      result_map["id"] = function_res["id"]
      result_map["error"] = nil
      *reply, _ = json.Marshal(result_map)

    default:
      *reply = []byte("This is not a valid function")
  }
  //*reply = "Have reached"
  return nil
}

func main(){
  global_map = make(map[string]interface{})
  global_map["mihirkelkar"] = "this is being returned from the map"

  //Instantiate a map
  listings := new(Listings)
  rpc.Register(listings)
  listener, error := net.Listen("tcp", ":1234")
  if error != nil{
    log.Fatal("Error occured while listening" + error.Error())
  }
  for{
    connection, err := listener.Accept()
    if err != nil{
      log.Fatal("Error while listening request")
    } else {
      log.Println("Connection to the client service established")
    }
    go rpc.ServeConn(connection)
  }
}
