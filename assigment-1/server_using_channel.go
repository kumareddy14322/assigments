// a simple server using go channel


package main

import (
"fmt"
"io"
"net/http"
"strconv"
"time"
)

func main() {
    c := AudioProcess()
    handleHello := makeHello(c)

    http.HandleFunc("/", handleHello)
    http.ListenAndServe(":8000", nil)
}

func makeHello(c chan string) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        for item := range c {                              // this loop runs when channel c is closed
            io.WriteString(w, item)
        }
    }
}

func AudioProcess() chan string {
    c := make(chan string)
    go func() {
        for i := 0; i <= 10; i++ {    // Iterate the audio file
            c <- strconv.Itoa(i)      // have my frame of samples, send to channel c
            time.Sleep(time.Second)
            fmt.Println("send ", i)   // logging
        }
        close(c) // done processing, close channel c
        }()
        return c
    }