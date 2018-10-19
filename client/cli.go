package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("need args")
        return
    }
    act := os.Args[1]
    var url string
    switch act {
    case "r":
        url = fmt.Sprintf("http://127.0.0.1:8090/act/rotation?face=%s&clockWise=%s", os.Args[2], os.Args[3])
    case "m":
        url = fmt.Sprintf("http://127.0.0.1:8090/act/missile?face=%s", os.Args[2])
    case "d":
        url = fmt.Sprintf("http://127.0.0.1:8090/act/dice?face=%s", os.Args[2])
    case "print":
        url = "http://127.0.0.1:8090/print"
    case "close":
        url = "http://127.0.0.1:8090/close"
    case "ping":
        url = "http://127.0.0.1:8090/ping"
    default:
        fmt.Println("act err")
        return
    }
    
    res, err := http.Get(url)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
    }

    defer res.Body.Close()

    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
    }

    fmt.Println(string(content))
}

