package main

import (
	"github.com/urfave/cli/v2"
	"github.com/gin-gonic/gin"
)

var runcmd = &cli.Command{
	Name:  "run",
	Usage: "run",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "port",
			Usage: "host  port the server api will listen on",
			Value: "8080",
		},
		&cli.StringFlag{
			Name:  "path",
			Usage: "file path to serve",
			Value: "./source",
		},

	},
	Description: "For more instructions see 'tinydoc --help'",
	Action: func(cctx *cli.Context) error {
		g := gin.Default()
		g.Static("/",cctx.String("path"))
		g.Run(":"+cctx.String("port"))
		return nil
	},
}