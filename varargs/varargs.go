package varargs

import (
  "fmt"
)

type Options struct {
  Par1 string
  Par2 int64
  Par3 interface{}
}

func PrintVarArgsUsingStruct(para *Options) {
  fmt.Print("\n\nPrintVarArgsUsingStruct: ", *para)
}

func PrintVarArgsUsingEmptyInterface(para ...interface{}) {
  fmt.Print("\n\nPrintVarArgsUsingEmptyInterface: \n")
  for idx, val := range para {
    switch val.(type) {
    case int:
      fmt.Print("The type is int. ")
    default:
      break
    }
    fmt.Print("idx: ", idx, " val: ", val, "\n")
  }
}