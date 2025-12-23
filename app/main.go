package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	var x string
	_,e:= fmt.Scan(&x)
	if e== nil{ fmt.Print(x,": command not found")
}else	{fmt.Print(e)}
}
