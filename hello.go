package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
  
  //  i++ is not an expression only a statement
  // also cannot ++i
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}
