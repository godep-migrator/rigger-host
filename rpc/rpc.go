package rpc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func Listen(port int) {
	log.Printf("Listening on port %v...", port)

	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(ProjectService), "")

	portString := fmt.Sprintf(":%v", port)

	http.Handle("/rpc", s)
	http.ListenAndServe(portString, nil)
}
