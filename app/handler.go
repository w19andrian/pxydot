package app

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
	"github.com/w19andrian/pxydot/config"
)

// Handler for handling incoming query and forward query
type Handler struct {
	client *dns.Client
	config config.Config
}

// NewHandler creates new handler
// to handle incoming query and forward query
func NewHandler(client *dns.Client, config config.Config) Handler {
	return Handler{
		client: client,
		config: config,
	}
}

// The ServeDNS act as a method to receive question
// from the client, forwards it to the upstream DNS, and write
// the answer back to the client
func (h Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	rHostPort := fmt.Sprintf("%s:%s", h.config.Upstream_Server, h.config.Upstream_Port)
	resp, _, err := h.client.Exchange(r, rHostPort)
	if err != nil {
		log.Printf("Error querying to upstream server: %v\n", err)
		return
	}

	w.WriteMsg(resp)
}
