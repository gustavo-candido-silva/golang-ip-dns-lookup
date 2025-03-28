package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna a aplicação de terminal pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de terminal"
	app.Usage = "Busca IPs e Nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de enderecos ip's na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServers,
		},
	}

	return app
}

func buscarServers(c *cli.Context) {
	var host string = c.String("host")

	names, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

func buscarIps(c *cli.Context) {
	var host string = c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
