package main

import (
  "C"
)

func main() {
}

//export EchoInt
func EchoInt(val int) int {
  return val
}

//export EchoBool
func EchoBool(val bool) bool {
  return val
}

//export EchoString
func EchoString(x *C.char) *C.char {
  input := C.GoString(x)
  return C.CString(input)
}

//export Echofloat
func Echofloat(x float64) float64 {
  return x
}
