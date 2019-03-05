package main

import "fmt"

type ItemType struct {
  Val interface{}
}

type Stack struct {
  maxItems, top int
  items []ItemType
}

type StackInterface interface {
  IsFull() bool
  IsEmpty() bool
  makeEmpty()
  Push(ItemType)
  Pop()
  Peak() ItemType
}

func (s Stack) IsFull() bool {
  if s.top == s.maxItems-1 {
    return true
  }
  return false
}

func (s Stack) IsEmpty() bool {
  if s.top == -1 {
    return true
  }
  return false
}

func (s *Stack) makeEmpty() {
  s.top = -1
}

func (s *Stack) Push(item ItemType) {
  s.top++
  s.items[s.top] = item
}

func (s *Stack) Pop() {
  s.top--
}

func (s *Stack) Peak() ItemType {
  return s.items[s.top]
}

func main() {
  var max int = 500
  s := &Stack{ maxItems: max, items: make([]ItemType,max), top: -1}
  item := ItemType{15}
  s.Push(item)
  item1 := s.Peak()
  fmt.Println(item1.Val)
  s.Pop()
}
