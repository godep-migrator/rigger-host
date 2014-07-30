package host

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func RPCListen(port int) {
	log.Printf("Listening on port %v...", port)

	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")

	portString := fmt.Sprintf(":%v", port)

	http.Handle("/rpc", s)
	http.ListenAndServe(portString, nil)
}
