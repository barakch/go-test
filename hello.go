package main // import "github.com/barakch/gotest"

import ("log"

        "net/http"

        "github.com/gorilla/mux")

// our main function
func main() {
        router := mux.NewRouter()
        log.Fatal(http.ListenAndServe(":8000", router))
}
