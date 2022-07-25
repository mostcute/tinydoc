package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime/debug"
)

func main() {

	local := []*cli.Command{
		versionCmd,
		runcmd,
	}
	app := &cli.App{
		Flags: []cli.Flag{

		},
		Commands: local,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "version",
	Flags: []cli.Flag{
	},
	Description: "For more instructions see 'tinydoc --help'",
	Action: func(cctx *cli.Context) error {
		fmt.Println (debug.ReadBuildInfo())
		return nil
	},
}