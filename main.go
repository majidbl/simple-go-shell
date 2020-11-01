package main

import (
     "os"
     "bufio"
     "fmt"
     "strings"
     "time"
     "reflect"
)
var inch = make(chan string)
var ticker = time.NewTicker(5 * time.Second)
func main() {
  for{
    //t := time.Now()
    go getinput()
    select{
      case msg := <-inch:
        fmt.Println(msg)
      case c := <-ticker.C:
      fmt.Println()
      fmt.Println(c)
    }
    
  }
}

func getinput(){
  
  
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    
    name, _ := reader.ReadString('\n')
    namesplit := strings.Split(name," ")
    str := strings.TrimRight(namesplit[len(namesplit)-1], "\n")
    if str == "e"{
      os.Exit(0)
    }
    inch <- namesplit[len(namesplit)-1]
    fmt.Println(reflect.TypeOf(namesplit[len(namesplit)-1]))
    
}