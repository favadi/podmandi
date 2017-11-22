package main

import (
	"errors"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/favadi/podmandi"
)

// dataFile returns the location of data file.
// TODO: implement this
func dataFile(appName string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dataDir := filepath.Join(usr.HomeDir, ".config", appName)
	if err = os.MkdirAll(dataDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dataDir, "data.json"), nil
}

func main() {
	app := cli.NewApp()
	app.Name = "podmandi"
	app.Usage = "A command line podcast manager"
	app.Version = "0.0.1"

	df, err := dataFile(app.Name)
	if err != nil {
		log.Fatal(err)
	}

	m, err := podmandi.NewManager(podmandi.WithDataFile(df))
	if err != nil {
		log.Fatal(err)
	}

	app.Commands = []cli.Command{
		{
			Name:   "add",
			Action: addAction(m),
		},
	}

	if err = app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func addAction(m *podmandi.Manager) func(context *cli.Context) error {
	return func(ctx *cli.Context) error {
		url := ctx.Args().First()
		if len(url) == 0 {
			return errors.New("missing URL")
		}

		return m.Add(url)
	}
}
