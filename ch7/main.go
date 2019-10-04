package main

import ("fmt")

func makeEvenGenerator() func() int {
  i := 0

  return func() (ret int) {
    ret = i
    i += 2
    return
  }
}

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func main() {
  // nextEven := makeEvenGenerator()
  // fmt.Println(nextEven()) // 0
  // fmt.Println(nextEven()) // 2
  // fmt.Println(nextEven()) // 4
  //
  // defer second()
  // first()

  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}
