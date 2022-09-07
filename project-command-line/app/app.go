package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generete retorna a aplicação de linha de comando pronta para ser executada
func Generete() *cli.App {
	app := cli.NewApp()
	app.Name = "Project Command Line"
	app.Usage = "Search IPs and Serve Names in WEB"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search Ips for adress in web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search Names Servers",
			Flags:  flags,
			Action: searchServes,
		},
	}

	return app
}

// Pegar o host e buscar o ip
func searchIps(c *cli.Context) {
	host := c.String("host")

	// Package net
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServes(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // Name Service
	if err != nil {
		log.Fatal(err)
	}

	for _, serve := range servers {
		fmt.Println(serve.Host)
	}
}
