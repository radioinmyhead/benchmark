package main

import (
	"github.com/radioinmyhead/benchmark"

	"github.com/urfave/cli"
)

var CmdDisk = cli.Command{
	Name:        "disk",
	Usage:       "disk `{dev}`",
	Description: `disk /dev/sdc`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:   "init",
			Hidden: false,
		},
		cli.StringFlag{
			Name: "grep",
		},
	},
	Action: diskbenchmark,
}

func diskbenchmark(c *cli.Context) error {
	if c.NArg() <= 0 {
		return cli.NewExitError("benchmark disk {dev}", 2)
	}

	disk, err := benchmark.NewDisk(c.Args().First())
	if err != nil {
		return err
	}
	/*
		if c.String("runtime") != "" {
			disk.Fio["--runtime"] = c.String("runtime")
		}
	*/

	if c.Bool("init") {
		if err := disk.Init(); err != nil {
			return err
		}
	}
	return disk.Benchmark()
}
