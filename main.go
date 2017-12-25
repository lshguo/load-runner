package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"math"
)

// cpu load runner
// GET http://localhost:8080/run

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/run").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	go func(){
		for i := 0; i < 65536; i++{
			_ = math.Sqrt(12.3456)
		}
	}()
	io.WriteString(resp, "running...\n")
}
