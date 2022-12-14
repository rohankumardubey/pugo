package cmd

import (
	"pugo/pkg/core/generator"
	"pugo/pkg/core/server"

	"github.com/urfave/cli/v2"
)

var (
	serverFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "port",
			Value: "18080",
			Usage: "port to listen",
		},
	}
)

// NewServer returns a new cli.Command for the server subcommand.
func NewServer() *cli.Command {
	flags := append(globalFlags, genFlags...)
	flags = append(flags, serverFlags...)
	cmd := &cli.Command{
		Name:        "server",
		Usage:       "server the site",
		Description: "start serve the site, auto rebuild when source files changed",
		Aliases:     []string{"serve"},
		Flags:       flags,
		Action: func(c *cli.Context) error {
			initGlobalFlags(c)

			opt := parseCliOption(c)
			opt.IsLocalServer = true

			generator.Generate(opt)

			s := server.New(server.ServerOption{
				Port: c.Int("port"),
				Dir:  opt.OutputDir,
			})
			s.Run()

			return nil
		},
	}
	return cmd
}
