package main

import (
	"fmt"
	// "bufio"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
 for{	
	fmt.Print("$ ")
	var x string
	_,e:= fmt.Scan(&x)
	if e!= nil{fmt.Print(e) 
		os.Exit(1)
	}
fmt.Print(x,": command not found\n")
}
}
