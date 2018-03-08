package main

import (
	"brain-data/benchmark/cmd/cmd"
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
		cmd.CmdDisk,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
