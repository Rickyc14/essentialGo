package myutils

import (
  "fmt"
  "strconv"
)

func Get_float(str string) float64 {
  for {
    var float_string string
    fmt.Print(str)
    fmt.Scanln(&float_string)
    float_number, error := strconv.ParseFloat(float_string, 64)
    if error == nil {
      return float_number
    }
    fmt.Println("Invalid input.")
  }
}

func Get_int(str string) int {
  for {
    var int_string string
    fmt.Print(str)
    fmt.Scanln(&int_string)
    int_number, error := strconv.Atoi(int_string)
    if error  == nil {
      return int_number
    }
    fmt.Println("Invalid input.")
  }
}

func Factorial(x int) int {
  if x == 1 || x == 0 {
    return 1
  }
  return x*Factorial(x-1)
}

func Power(x float64, n int) float64 {
  if n == 0 {
    return 1
  }
  return x*Power(x, n-1)
}
