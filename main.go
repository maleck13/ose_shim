package main

import (
	"github.com/codegangsta/cli"
	"github.com/maleck13/ose_shim/cmd"
	"os"
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

