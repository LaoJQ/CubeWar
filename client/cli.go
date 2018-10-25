package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

const BaseUrl = "http://127.0.0.1:8090/"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("need args")
        return
    }
    act := os.Args[1]
    var url string
    switch act {
    case "r":
        url = fmt.Sprintf(BaseUrl+"act/rotation?face=%s&target=%s&clockWise=%s", os.Args[2], os.Args[3], os.Args[4])
    case "m":
        url = fmt.Sprintf(BaseUrl+"act/missile?face=%s", os.Args[2])
    case "d":
        url = fmt.Sprintf(BaseUrl+"act/dice?face=%s", os.Args[2])
    case "b":
        url = fmt.Sprintf(BaseUrl+"act/blood?face=%s", os.Args[2])
    case "print":
        url = BaseUrl+"print"
    case "close":
        url = BaseUrl+"close"
    case "ping":
        url = BaseUrl+"ping"
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

