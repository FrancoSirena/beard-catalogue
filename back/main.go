package main

import (
    "bytes"
    "flag"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "github.com/rs/cors"
)

type Resp struct {
    Something string
}

func res(w http.ResponseWriter, r *http.Request) {
    var vv Resp = Resp{ "wrong" }
    fmt.Println(r)
    j, _ := json.Marshal(vv)
    w.Header().Set("Content-Type", "application/json")
    io.Copy(w, bytes.NewReader(j))

}
func main() {
    var port int
    flag.IntVar(&port, "port", 80, "The port to listen on")
    flag.Parse()

    ss := http.NewServeMux()
    c:=cors.New(cors.Options{
      AllowedOrigins: []string{"http://localhost:8081"},
      AllowCredentials: true,
    })

    ss.HandleFunc("/api/v1/try", res)
    handler := c.Handler(ss)
    log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
