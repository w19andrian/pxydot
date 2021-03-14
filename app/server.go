package app

import (
	"log"

	"github.com/miekg/dns"
	"github.com/w19andrian/pxydot/config"
)

var tsvr *dns.Server
var usvr *dns.Server

func RunServer(c config.Config) {

	if !c.TCP_Enabled && !c.UDP_Enabled {
		log.Fatalf("Need at least one enabled protocol. Exitting...")
	}
	if c.TCP_Enabled {
		tsvr = &dns.Server{Addr: c.Listen_Addr, Net: "tcp"}
		go tsvr.ListenAndServe()
		log.Printf("Listening on %s %s", c.Listen_Addr, "tcp")
	}
	if c.UDP_Enabled {
		usvr = &dns.Server{Addr: c.Listen_Addr, Net: "udp"}
		go usvr.ListenAndServe()
		log.Printf("Listening on %s %s", c.Listen_Addr, "udp")
	}
	log.Printf("Upstream server: %s:%s", c.Upstream_Server, c.Upstream_Port)
}

func Shutdown(svr *dns.Server) {
	if svr == nil {
		return
	}
	if err := svr.Shutdown(); err != nil {
		log.Printf("Shutdown failed: %v", err)
	}
}
