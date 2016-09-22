package main

import (
    "fmt"
    "golang.org/x/net/websocket"
    //"io"
    "net/http"
    "strings"
    "time"
)

func echoHandler(ws *websocket.Conn) {
    fmt.Println("Inside web socket handler:", time.Now())
    //io.Copy(ws, ws)
///*
    var reply string
    var err error
    for {
        fmt.Println("Inside loop:", time.Now())
        err = websocket.Message.Receive(ws, &reply);
        if err != nil {
            fmt.Println("Can't receive")
            break
        }

        fmt.Println("Received from client: " + reply)

        msg := strings.ToUpper(reply)
        fmt.Println("Sent to client: " + msg)
        err = websocket.Message.Send(ws, msg);
        if err != nil {
            fmt.Println("Can't send")
            break
        }
    } // for
//*/
}

func main() {
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))
    var address string = "0.0.0.0:8080";
    fmt.Println("Listening on ", address);
    err := http.ListenAndServe(address, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}