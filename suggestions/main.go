package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// DictPath is the path to the dictionary, provided via a mandatory flag.
var DictPath string

// CorrectionsCommand displays a possible correction of the given name if any.
func CorrectionsCommand(c *cli.Context) error {
	name := c.Args().Get(0)
	if !validName(name) {
		return cli.NewExitError("must provide a valid name", 1)
	}
	if !validDict() {
		return cli.NewExitError("must provide a valid dictionary", 1)
	}

	sugs, ok := findSuggestionsFor(name)
	if ok {
		log.Printf("'%s' looks like a valid name", name)
	} else if len(sugs) > 0 {
		printSuggestions(sugs)
	} else {
		log.Printf("we found no suggestions for '%s' :(", name)
	}
	return nil
}

func printSuggestions(names []Name) {
	total := 0
	for _, n := range names {
		total = total + n.freq
	}

	log.Printf("perhaps you meant any of")
	for _, n := range names {
		percent := float64(n.freq) / float64(total)
		log.Printf("\t * %s (%f%%)", n.word, percent)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "name_suggestions"
	app.Usage = "suggests spelling errors based on a given name dictionary"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dict, d",
			Usage:       "load names dictionary from `FILE`",
			Destination: &DictPath,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "suggest",
			Aliases: []string{"s"},
			Usage:   "display a possible correction of the given name if any",
			Action:  CorrectionsCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
