package main

import (
  "fmt"
  common "golang-book/ch11/common"
)

func main() {
  x:= []float64{1.0,2.0}

  fmt.Println(common.Average(x))
}
