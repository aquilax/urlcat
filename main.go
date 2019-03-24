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
		stopOnError: true,
	}
}

func processStream(c *config, r reporter) error {
	scanner := bufio.NewScanner(c.from)
	for scanner.Scan() {
		line := scanner.Text()
		url, err := url.Parse(line)
		if err != nil && c.stopOnError {
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
		Name:        "urlcat",
		Usage:       "url processing tool. Reads line separated ",
		Description: "Reads Line separated URL's from stdin and returns the requested segment",
		Version:     "1.0.0",
		Authors: []cli.Author{
			cli.Author{
				Name:  "aquilax",
				Email: "aquilax@gmail.com",
			},
		},
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
				Usage:   "Returns the host",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnHost)
				},
			},
			{
				Name:    "scheme",
				Aliases: []string{"s"},
				Usage:   "Returns the URL scheme",
				Action: func(c *cli.Context) error {
					return processStream(newConfig(), returnProtocol)
				},
			},
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Returns the query string",
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
