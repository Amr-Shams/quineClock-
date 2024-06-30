package main 
import (
"fmt")

func main(){
    src:= "package main \nimport (\n\"fmt\")\n\nfunc main(){\n    src:= \"?\"\n    for _,c:= range src{\n        if c == 63{\n            for _,ch:= range src{\n                switch ch{\n                    case '\\n':\n                        fmt.Printf(\"\\\\n\")\n                        break\n                    case '\"':\n                        fmt.Printf(\"\\\\\\\"\")\n                        break\n                    case '\\\\':\n                        fmt.Printf(\"\\\\\\\\\")\n                        break\n                    default:\n                        fmt.Printf(\"%c\",ch)\n                }\n            }\n        }else{\n            fmt.Printf(\"%c\",c)\n        }\n    }\n}\n"
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
