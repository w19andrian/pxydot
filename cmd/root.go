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

func Start() {
	sigsend := make(chan os.Signal, 1)
	signal.Notify(sigsend, syscall.SIGTERM)
	signal.Notify(sigsend, syscall.SIGINT)

	config.LoadConfig()

	c := new(dns.Client)
	c.Net = "tcp-tls"

	h := app.NewHandler(c, config.AppConfig)

	app.RunServer(config.AppConfig)
	dns.Handle(".", h)

	sigrec := <-sigsend
	log.Printf("Received signal %s", sigrec.String())

}
