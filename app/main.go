package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	// "unicode"
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
	// fmt.Printf("%q\n",x)    // raw input 

	in:= parsed_args(x)[0]
	args := parsed_args(x)[1:]
	// fmt.Println(args)

	// fmt.Printf("in: %q \n arg: %q\n",in,args) // for out of input and arguments
	if e!= nil{
		fmt.Print(e) 
		os.Exit(1)
	}
		var rd_arg string ;
		var redirect, append_cmd, rd_err  bool

		for i,r := range args{
		if r==">"||r=="1>"||r=="2>"||r=="1>>"||r==">>" {
			rd_arg =args[i+1]
			args = args[:i]  // to make run only the part of commnad, before the special redirect symbols
			// fmt.Println(args)
			redirect=true
			if r=="2>"{
			rd_err=true
			redirect=false
			}
			if r=="1>>"||r==">>" {
				append_cmd =true
			}
			
		}
		
	}
	builtin_cmds := []string{"echo","type","exit","cd","pwd"} 

	switch in {
	
	case "echo":

		if redirect {
			output :=strings.Join(args," ")+"\n"
				if append_cmd {
					content,_:=os.ReadFile(rd_arg)
					output = string(content)+output+"\n"
					append_cmd = false
				}
			os.WriteFile(rd_arg,[]byte(output),0666)
			redirect=false
		}else{
		fmt.Println(strings.Join(args," "))
			}
		if rd_err {
				os.WriteFile(rd_arg,[]byte(""),0666)
				rd_err=false
				continue;
			}
	case "type":
		 for _,cmd:= range args  {
			
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
			  var stderr strings.Builder
				proc.Stderr=&stderr
				proc.Stdout=&out
				err:= proc.Run()
			
				if rd_err {
				os.WriteFile(rd_arg,[]byte(stderr.String()),0666)
				rd_err=false
					if out.String()!=""{
						fmt.Print(out.String())
					}
					continue;
				}
				if err!=nil { 
					fmt.Print(stderr.String())
				}
			    if redirect {
					output := out.String()
					if append_cmd {
						content,_:=os.ReadFile(rd_arg)
						output =string(content)+ output
						append_cmd = false
					}
					os.WriteFile(rd_arg,[]byte(output),0666)
					redirect=false
				}else{
				fmt.Print(out.String())
				}

			}	else{
				fmt.Println(in+": command not found")
			}
	}
	}
}
// 	TODO: MAKE REDIRECTION IN PARSE FUNCTION