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
	in :=strings.TrimSpace(strings.Split(x," ")[0])
	a:= strings.SplitAfter(x," ")[1:]
	arg:= strings.Trim(strings.Join(a, "")," ")
	// fmt.Printf("in: %q \n arg: %q\n",in,a)
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	switch in {
	case "echo":
		 fmt.Print(arg)
	case "exit":
		 os.Exit(0)
	default:
fmt.Print(in,": command not found\n")
	}
}
}
