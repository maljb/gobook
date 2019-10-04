package main

import (
  "fmt"
)

func main(){
  fmt.Println(`1
2
3
4
5
6
7
8
9
10`)

  for i := 1; i <= 10; i++ {

    if i % 2 == 0 {
     // even
     fmt.Println("i", "even")
    } else {
     // odd
     fmt.Println(i, "odd")
    }
  }

  // fmt.Println(i)
}
