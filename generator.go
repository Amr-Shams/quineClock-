package main 
import (
    "fmt"
    "os"
    "bufio"
)

func main(){
    if len(os.Args)<2 {
        fmt.Println("Usage: go run generator.go -- <template>")
        os.Exit(1)
    }
    template:=os.Args[2]
    // read the file
    file, err := os.Open(template)
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    defer file.Close()
    // read the entire file and store it in a string
    scanner := bufio.NewScanner(file)
    var content string
    for scanner.Scan() {
        content+=scanner.Text()+"\n"
    }
    for _,char:=range content {
        switch char {
        case '?':
            for _,c:=range content{
                switch c {
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
                        fmt.Printf("%c",c)
                }
            }
        break 
        default:
            fmt.Printf("%c",char)
         }
}
}
