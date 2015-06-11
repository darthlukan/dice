package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func Dfour(qty int64) {}

func Dsix(qty int64) {}

func Deight(qty int64) {}

func Dten(qty int64) {}

func Dtwelve(qty int64) {}

func Dtwenty(qty int64) {}

func Roll(sides, qty int64) {}

func main() {
	app := cli.NewApp()
	app.Name = "dice"
	app.Version = "0.0.1"
	app.Usage = "Roll the dice, I dare ya'!"
	app.Authors = []cli.Author{cli.Author{
		Name:  "Brian Tomlinson",
		Email: "darthlukan@gmail.com",
	}}
	app.Run(os.Args)
}
