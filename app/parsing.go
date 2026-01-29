package main

import(
	"strings"
	"unicode"
)

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
		
		toggle=!toggle

		if t2 && !(c=='"'||c=='\\'||c=='$'||c=='\n'||c=='`'){
			sb.WriteRune('\\')	
			sb.WriteRune(c)

			continue
		}
		sb.WriteRune(c)
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
