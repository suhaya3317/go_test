package main

import "fmt"

type Sample struct {
  X, Y int
}

//こちらがメソッド
func (s *Sample) renderValue() {
  fmt.Println(s.X)
  fmt.Println(s.Y)
}

func main()  {
  s := Sample{X: 1, Y: 2}
  s.renderValue()
}
