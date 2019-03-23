package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"

	"gopkg.in/urfave/cli.v1"
)

type reporter = func(u *url.URL) string

type config struct {
	from        io.Reader
	to          io.Writer
	stopOnError bool
}

func newConfig() *config {
	return &config{
		from:        os.Stdin,
		to:          os.Stdout,
		stopOnError: false,
	}
}

func processStream(c *config, r reporter) error {
	scanner := bufio.NewScanner(c.from)
	for scanner.Scan() {
		line := scanner.Text()
		url, err := url.Parse(line)
		if err != nil {
			return err
		}
		fmt.Fprintln(c.to, r(url))
	}
	return nil
}

func returnFilename(u *url.URL) string { return path.Base(u.Path) }
func returnProtocol(u *url.URL) string { return u.Scheme }
func returnQuery(u *url.URL) string    { return u.RawPath }
func returnHost(u *url.URL) string     { return u.Host }

func main() {
	app := &cli.App{
		Name:  "urlcat",
		Usage: "url processing tool",
		Commands: []cli.Command{
			{
				Name:    "filename",
				Aliases: []string{"f"},
				Usage:   "returns the filename (default)",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnFilename)
				},
			},
			{
				Name:    "host",
				Aliases: []string{"h"},
				Usage:   "returns the host",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnHost)
				},
			},
			{
				Name:    "scheme",
				Aliases: []string{"s"},
				Usage:   "returns the scheme",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnProtocol)
				},
			},
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "returns the query",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnQuery)
				},
			},
		},
		Action: func(c *cli.Context) error {
			return processStream(newConfig(), returnFilename)
		},
	}

	app.Run(os.Args)
}
