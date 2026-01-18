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
	"unicode"
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
var toggle bool =false
func quoted_args( c rune ) (bool){
	if(c=='\''){
		toggle = !toggle
		return true;
	}else	if (toggle==true){
		return false
	} 
return unicode.IsSpace(c);
}

func parsed_args( raw string) string{
	var sb strings.Builder
		count:=0
		for _,c:= range raw{
			
	if(c=='\''){
		toggle = !toggle
		count=0
		continue;
	}else	if (toggle){
		sb.WriteRune(c)
		continue
	} else if (unicode.IsSpace(c)&& !toggle && count==0){
		sb.WriteRune(c)
	}else{
		sb.WriteRune(c)
		count=0
	}
		}
		return sb.String();
}


func main() {
	
	for{	
	fmt.Print("$ ")

	x,e:= bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Printf("%q\n",x)    // raw input 
	in :=strings.TrimSpace(strings.Split(x," ")[0])
	args :=strings.FieldsFunc(x,quoted_args)[1:]
	   
	raw:=strings.Join(strings.Split(x," ")[1:]," ")
		
// fmt.Println(sb.String())           //string bulder output
	// echo_out:= strings.Trim(strings.Join(args, " "),"\n")
	// fmt.Printf("in: %q \n arg: %q\n",in,args) // for out of input and arguments
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	builtin_cmds := []string{"echo","type","exit","cd","pwd"} 

	switch in {
	
	case "echo":
		 fmt.Println(strings.TrimSuffix(parsed_args(raw),"\n"))
	
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
	
	case "cd":
		if len(args)==1 {	
			os.Chdir(args[0])
		}else if len(args)>1{
		fmt.Println("Too many arguments")
	}
	// case "pwd":
			
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

