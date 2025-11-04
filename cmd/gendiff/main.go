package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "code",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			if c.NArg() < 2 {
				return cli.Exit("Error: two file paths are required\nUsage: code [options] <file1> <file2>", 2)
			}
			file1 := c.Args().Get(0)
			file2 := c.Args().Get(1)
			format := c.String("format")

			out, err := code.GenDiff(file1, file2, format)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Println(out)
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
