package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"time"
)

func randomInt64(n int64) int64 {
	// rand.Int63n panics if n == 0
	if n == 0 {
		n = 1
	}
	return rand.Int63n(n) + 1 // Never want to return zero
}

func sidesAreValid(sides int64) bool {
	validSides := []int64{4, 6, 8, 10, 12, 20}
	for _, vs := range validSides {
		if vs == sides {
			return true
		}
	}
	return false
}

func Roll(sides, qty int64) int64 {
	var val int64
	for i := int64(1); i <= qty; i++ {
		val += randomInt64(sides)
	}
	return val
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	app := cli.NewApp()
	app.Name = "dice"
	app.Version = "0.0.1"
	app.Usage = "Roll the dice, I dare ya'!"
	app.Authors = []cli.Author{cli.Author{
		Name:  "Brian Tomlinson",
		Email: "darthlukan@gmail.com",
	}}
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "sides, s",
			Value: 0,
			Usage: "The number of sides of the desired die.",
		},
		cli.IntFlag{
			Name:  "qty, q",
			Value: 0,
			Usage: "The number of die to roll.",
		},
	}
	app.Action = func(c *cli.Context) {
		sides := int64(c.Int("sides"))
		qty := int64(c.Int("qty"))

		if sidesAreValid(sides) && qty >= 1 {
			result := Roll(sides, qty)
			fmt.Printf("%vd%v roll result: %v\n", qty, sides, result)
		} else {
			fmt.Printf("Please see 'dice --help' for usage assistance.\n")
		}
	}
	app.Run(os.Args)
}
