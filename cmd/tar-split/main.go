// go:generate git tag | tail -1
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tar-split"
	app.Usage = "tar assembly and disassembly utility"
	app.Version = "0.9.2"
	app.Author = "Vincent Batts"
	app.Email = "vbatts@hashbangbash.com"
	app.Action = cli.ShowAppHelp
	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stderr)
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "debug output",
			// defaults to false
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "disasm",
			Aliases: []string{"d"},
			Usage:   "disassemble the input tar stream",
			Action:  CommandDisasm,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output",
					Value: "tar-data.json.gz",
					Usage: "output of disassembled tar stream",
				},
			},
		},
		{
			Name:    "asm",
			Aliases: []string{"a"},
			Usage:   "assemble tar stream",
			Action:  CommandAsm,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input",
					Value: "tar-data.json.gz",
					Usage: "input of disassembled tar stream",
				},
				cli.StringFlag{
					Name:  "output",
					Value: "-",
					Usage: "reassembled tar archive",
				},
				cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "relative path of extracted tar",
				},
			},
		},
		{
			Name:   "checksize",
			Usage:  "displays size estimates for metadata storage of a Tar archive",
			Action: CommandChecksize,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "work",
					Usage: "do not delete the working directory",
					// defaults to false
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
