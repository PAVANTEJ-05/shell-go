package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// path for executables
func pathOf(cmd string) (string,bool){
		p:= os.Getenv("PATH")
		path:= strings.SplitSeq(p,string(os.PathListSeparator))

		for dir:= range path{

			fullpath:= filepath.Join(dir,cmd)
			exist,_:=filepath.Match( filepath.Join(dir,"/*"), fullpath)

			f,err:= os.Stat(fullpath)

			if err==nil && exist && (f.Mode().Perm()&0111 !=0){
				return fullpath,true
			}
		}
		return "",false
}

func main() {

	for{	
	fmt.Print("$ ")

	x,e:= bufio.NewReader(os.Stdin).ReadString('\n')
	in :=strings.TrimSpace(strings.Split(x," ")[0])
	args :=strings.Fields(x)[1:]
	echo_arg:= strings.Trim(strings.Join(args, " "),"\"\n")
	// fmt.Printf("in: %q \n arg: %q\n",in,arg)
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	builtin_cmds := []string{"echo","type","exit"} 

	switch in {
	
	case "echo":
		 fmt.Println(echo_arg)
	
	case "type":
		 for _,cmd:= range args  {
			
			cmd:=strings.TrimSpace(cmd)
			// for executables path 
			path,_:=pathOf(cmd)
			
			if slices.Contains(builtin_cmds,cmd) {
			  fmt.Println(strings.TrimSpace(cmd),"is a shell builtin")
				continue;
			}else if (path!="") {
				fmt.Println(cmd,"is",path)
			 }else{
				fmt.Println(cmd+": not found")
			 }
		}
	
	case "exit":
		 os.Exit(0)
	
	default:
		_,exist:=pathOf(in)
		if exist {
			  proc:= exec.Command(in,args...)
			  var out strings.Builder

				proc.Stdout=&out
				err:= proc.Run()
				if err!=nil{ log.Fatal(err)}
			    
				fmt.Print(out.String())

			}	else{
				fmt.Print(in,": command not found\n")
			}
	}
	}
}

