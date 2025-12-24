package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	arg:= strings.Trim(strings.Join(a, "")," \"\n")
	// fmt.Printf("in: %q \n arg: %q\n",in,arg)
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	commands := []string{"echo","type","exit"} 
	switch in {
	case "echo":
		 fmt.Println(arg)
	case "type":
		 for _,cmd:= range a  {
			cmd:=strings.TrimSpace(cmd)
			if slices.Contains(commands,cmd) {
			  fmt.Println(strings.TrimSpace(cmd),"is a shell builtin")

			}else{
				fmt.Print(cmd,": not found\n") }
		 }
	case "exit":
		 os.Exit(0)
	default:
fmt.Print(in,": command not found\n")

	}
}
}
