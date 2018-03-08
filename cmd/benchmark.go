package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const appVer = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "benchmark"
	app.Usage = ""
	app.Version = appVer
	app.Commands = []cli.Command{
		CmdDisk,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
