package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "time"
  "strconv"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/core/chaincode/lib/cid"
  sc "github.com/hyperledger/fabric/protos/peer"
)

// Represents our smartcontract.
type SmartContract struct {
}

type Person struct {
  Id string `json:"id"`
  Class string `json:"class"`
  Name string `json:"name"`
  Email string `json:"email"`
}

type Art struct {
  Id string `json:"id"`
  Class string `json:"class"`
  Name string `json:"name"`
  Description string `json:"description"`
  Artist string `json:"artist"`
  Owner string `json:"owner"`
  CreatedAt time.Time `json:"created_at"`
}

// Init function... called during chaincode instantiation to initialize any data.
// Upgrade also calls this function to reset or to migrate data
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Responce {
  return shim.Success(nil)
}

// Invoke function....it is like a POST-Call which is used on web.
func (s *SmartContract) CreateUser(APIstub shim.ChaincodeStubInterface) sc.Responce {
  Id := "user-"+utils.RandStringBytes(32)
  var user = Person{Class:"Person", Id: Id, Name: args[0], Email: args[1]}
  UserBytes, _ := json.Marshal(user)
  APIstub.PutState(Id, UserBytes)
  return shim.Success(nil)
}


// Query... to query the ledger, filter results based on particular logic and so on...
// returns the string which will be an array of results
func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting UserID")
  }
  UserBytes, err := APIstub.GetState(args[0])
  if err != nil {
    return shim.Error(err.Error())
  }
  return shim.Success(UserBytes)
}
