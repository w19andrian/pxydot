package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/miekg/dns"
	"github.com/w19andrian/pxydot/app"
	"github.com/w19andrian/pxydot/config"
)

var tcp_svr *dns.Server
var udp_svr *dns.Server

func init() {

	// Initiate configuration file to be loaded
	config.LoadConfig()
}

func Start() {

	// Create channel to accept signal SIGTERM and SIGINT
	// The app will fall into shutdown sequence if either
	// one of the signal is received
	notifSignal := make(chan os.Signal, 1)
	signal.Notify(notifSignal, syscall.SIGTERM)
	signal.Notify(notifSignal, syscall.SIGINT)

	// Initiate new client
	client := new(dns.Client)
	client.Net = "tcp-tls"

	// Initiate upstream servers and preparing them for
	// load balancing
	lb := app.InitUpstreamServers(config.AppConfig)
	handler := app.NewHandler(client, lb)

	// Running the servers
	func(c *config.Config) {
		if !c.TCP_Enabled && !c.UDP_Enabled {
			log.Fatal("At least one of the protocol should be enabled (TCP || UDP). Exitting...")
		}
		if c.TCP_Enabled {
			tcp_svr := &dns.Server{Addr: c.Listen_Addr, Net: "tcp"}
			go app.RunServer(tcp_svr)
			log.Printf("Server is Running and Listening on TCP/%s", c.Listen_Addr)
		}
		if c.UDP_Enabled {
			svr := &dns.Server{Addr: c.Listen_Addr, Net: "udp"}
			go app.RunServer(svr)
			log.Printf("Server is Running and Listening on UDP/%s", c.Listen_Addr)
		}
	}(config.AppConfig)

	dns.Handle(".", handler)

	sigrcv := <-notifSignal
	log.Printf("Received signal %s", sigrcv.String())

	// Shutdown sequences that will be invoked once SIGTERM or SIGINT
	// is received on the receiver channel
	func(c *config.Config) {
		if c.TCP_Enabled {
			log.Println("Shutting down TCP Server")
			app.Shutdown(tcp_svr)
		}
		if c.TCP_Enabled {
			log.Println("Shutting down UDP Server")
			app.Shutdown(udp_svr)
		}
	}(config.AppConfig)

}
