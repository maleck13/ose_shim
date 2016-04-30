package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/maleck13/ose_shim/cmd"
)

func main() {

	app := cli.NewApp()
	app.Name = "ose_shim"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}
