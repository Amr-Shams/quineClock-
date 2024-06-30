package main 
import (
"fmt")

func mainF(){
    src:= "?"
    for _,c:= range src{
        if c == 63{
            for _,ch:= range src{
                switch ch{
                    case '\n':
                        fmt.Printf("\\n")
                        break
                    case '"':
                        fmt.Printf("\\\"")
                        break
                    case '\\':
                        fmt.Printf("\\\\")
                        break
                    default:
                        fmt.Printf("%c",ch)
                }
            }
        }else{
            fmt.Printf("%c",c)
        }
    }
}
