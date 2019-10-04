package main

import (
  "fmt"
)

func main(){
  // var x string
  // x = "first"
  // fmt.Println(x)
  // x = x + "second"
  // fmt.Println(x)

  // var x string = "hello"
  // var y string = "hello"
  // fmt.Println(x == y)

  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := input * 2
  fmt.Println(output)

}
