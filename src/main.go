package main

import (
    "bufio"
//    "bytes"
    "fmt"
 //   "log"
    "os"
//    "os/exec"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("oh noes")
        return
    }

    filename := os.Args[1]
    fmt.Println("script:", filename)
    fmt.Println("-----------")
    fmt.Println("")

    // opening file & setting up cmd variable
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("ERROR")
    }
    defer file.Close()


    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        switch {
          
        case strings.HasPrefix(line, "cd"):
          path := strings.TrimSpace(strings.TrimPrefix(line, "cd"))
          os.Chdir(path)


        case line == "ls":
          ls, err := os.ReadDir(".")
          if err != nil {
            fmt.Println("I also have no idea why ls didn't work - G0o53") 
          }
          fmt.Println(ls)

        case line == "pwd":
          pwd, err := os.Getwd()
          if err != nil {
            fmt.Println("huh? it broke... idk check ur code")
          }
          fmt.Println(pwd)


        default:
          fmt.Println("eror")
          return

        case line == "":
          continue
        }

    // Check if the loop stopped because of an error instead of EOF
    if err := scanner.Err(); err != nil {
        fmt.Println("err;")
    }
  }
}
