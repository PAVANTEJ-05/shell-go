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

func parsed_args( raw string) []string{
	var t1,t2,toggle bool =false,false,false
	var sb strings.Builder
	var argv []string;
	count:=0
		for i,c:= range raw{
	
	if i==len(raw)-1 {
			argv = append(argv, sb.String())
			break;
	}else if toggle {

		sb.WriteRune(c)
		toggle=!toggle

		if(c=='\'' && t2){
		t1 = !t1
		count=0
		continue;

	}
		continue;

	}else if (c=='\\'&& !t1 ){
		toggle =!toggle
		continue

	}else if(c=='"' && !t1){
		t2 = !t2
		count=0
		continue;

	}else	if(c=='\'' && !t2){
		t1 = !t1
		count=0
		continue;

	}else	if (t1 || t2){
		sb.WriteRune(c)
		continue

	} else if (unicode.IsSpace(c)&& !t1 && !t2 && count==0){
		// sb.WriteRune(c)
		argv = append(argv,sb.String())
		sb.Reset();
		count++ ;	
		continue;
	}else if(!unicode.IsSpace(c)){
		sb.WriteRune(c)
		count=0
	}
		}
		return argv;
}


func main() {
	
	for{	
	fmt.Print("$ ")

	x,e:= bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Printf("%q\n",x)    // raw input 

	in:= parsed_args(x)[0]
	args := parsed_args(x)[1:]
	// fmt.Println(args)

	// fmt.Printf("in: %q \n arg: %q\n",in,args) // for out of input and arguments
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
	builtin_cmds := []string{"echo","type","exit","cd","pwd"} 

	switch in {
	
	case "echo":

		fmt.Println(strings.Join(args," "))
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

