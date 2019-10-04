package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
}

type Circle struct {
  x, y, r float64
}

type Square struct {
  l float64
}


func circleArea(c Circle) float64 {
  return math.Pi * c.r*c.r
}

func circleArea2(c *Circle) float64 {
  return math.Pi * c.r*c.r
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func (s *Square) area() float64 {
  return s.l*s.l
}

func test(s Shape) float64 {
  return s.area()
}

func main(){
  c := Circle{x: 0, y: 0, r: 5}
  s := Square{l: 2}

  fmt.Println(circleArea(c))
  fmt.Println(circleArea2(&c))
  fmt.Println(c.area())
  fmt.Println(test(&c))
  fmt.Println(test(&s))

}
