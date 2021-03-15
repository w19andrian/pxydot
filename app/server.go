package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/miekg/dns"
)

// type Handler for building query handler that implements dns.HandleFunc
type Handler struct {
	client   *dns.Client
	upstream *PoolUpstream
}

// RunServer will run the server to listen and serve connections
func RunServer(svr *dns.Server) error {
	return svr.ListenAndServe()
}

// NewHandler creates new handler for dns.HandlerFunc implementation
// to handle incoming query and forward query
func NewHandler(client *dns.Client, upst *PoolUpstream) Handler {
	return Handler{
		client:   client,
		upstream: upst,
	}
}

// The ServeDNS act as a method to receive question
// from the client, forwards it to the upstream DNS, and write
// the answer back to the client
func (h Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	next := h.upstream.GetNextUpstream()
	addr := next.Addr

	q := ""
	for _, v := range r.Question {
		q += strings.Split(v.String(), "\t")[0]
	}
	log.Printf("received query: %s", q)

	rHostPort := fmt.Sprintf("%s:%s", addr, "853")
	resp, _, err := h.client.Exchange(r, rHostPort)
	if err != nil {
		log.Printf("Error querying to upstream server: %v\n", err)
		return
	}
	log.Printf("Querying to: %s", addr)

	w.WriteMsg(resp)
}

// Shutdown checks if the server is active before actually
// invoking the Shutdown function
func Shutdown(svr *dns.Server) {
	if svr == nil {
		return
	}
	if err := svr.Shutdown(); err != nil {
		log.Printf("Shutdown failed: %v", err)
	}
}
