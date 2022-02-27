package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "strings"
)

/*import (
        "fmt"
        "log"
        "net/http"
        "os"
        "strings"
)

func index(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1> Welcome to my page</h1>"))

        for k, v := range r.Header {
                for _, vv := range v {
                        w.Header().Set(k,vv)
                }
        }

        os.Setenv("VERSION","1.0")
        version := os.Getenv("VERSION")
        w.Header().Set("VERSION",version)

        clientIP := getCurrentIP(r)
        httpCode := http.StatusOK
        log.Printf("clientip: %s, status code: %d \n",clientIP,httpCode)
}

func getCurrentIP(r *http.Request) string {
        ip := r.Header.Get("X-REAL-IP")

        if ip == "" {
                ip = strings.Split(r.RemoteAddr,":")[0]
        }
        return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"I am working")
}

func main() {
        mux := http.NewServeMux()

        mux.HandleFunc("/",index)
        if err := http.ListenAndServe(":8080",mux); err != nil {
                log.Fatalf("start failed,%s \n", err.Error())
        }
        mux.HandleFunc("/healthz",healthz)

}*/


func index(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1>Welcome to my page</h1>"))

        for k,v := range r.Header {
                for _,vv := range v {
                        w.Header().Set(k,vv)
                }
        }

        os.Setenv("VERSION","0.0.1")
        version := os.Getenv("VERSION")
        w.Header().Set("VERSION",version)

        clientIP := getCurrentIP(r)
        httpCode := http.StatusOK
        log.Printf("clientip: %s, status code: %s",clientIP,httpCode)
}

func getCurrentIP(r *http.Request) string {
        ip := r.Header.Get("X-REAL-IP")

        if ip == "" {
                ip = strings.Split(r.RemoteAddr,":")[0]
        }
        return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"healthz\n")
}

func main() {
        mux := http.NewServeMux()

        mux.HandleFunc("/",index)
        mux.HandleFunc("/healthz",healthz)
        if err := http.ListenAndServe(":8080",mux); err != nil {
                log.Fatal("start server failed, %s \n",err.Error())
        }
}
