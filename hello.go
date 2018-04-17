package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
  
  //  i++ is not an expression only a statement
  // also cannot ++i
  i := 1
  // apparently have  to include empty semicolon for init
  for i < len(os.Args) {
    s += sep + os.Args[i]
    sep = " "
    i++
  }
  fmt.Println(s)
}
