package main

import "fmt"

func binary_search(list []int, value int) bool {
  middle := len(list) / 2
  midval := list[middle]
  if midval == value {
    return true
  } else if len(list) == 1{
    return false
  } else if value < midval {
    return binary_search(list[:middle], value)
  } else if value > midval {
    return binary_search(list[middle:], value)
  } else {
    return false
  }
}

func main() {
  evenList := []int{1, 2, 3, 4}
  oddList := []int{1, 2, 3, 4, 5}
  fmt.Println(binary_search(oddList, 1))
  fmt.Println(binary_search(oddList, 10))
  fmt.Println(binary_search(evenList, 4))
  fmt.Println(binary_search(evenList, 10))
}
