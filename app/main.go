package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {

	for{	
	fmt.Print("$ ")

	x,e:= bufio.NewReader(os.Stdin).ReadString('\n')
	in :=strings.Split(x," ")[0]
	arg,_ :=strings.CutPrefix(x,in+" ")
		// fmt.Printf("in: %q \n arg: %q\n",in,arg)
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	switch in {
	case "echo":
			// i,_:= bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Print(arg)
	case "exit\n": os.Exit(0)
	default:
fmt.Print(in,": command not found\n")
	}
}
}
