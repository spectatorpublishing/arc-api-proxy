package main

import (
	"arc-api-proxy/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"os"
	"runtime"
)

func main() {
	StartServer()

}

func StartServer() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := NewRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + os.Getenv("PORT"))
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
