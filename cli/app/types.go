package app

import (
	"github.com/urfave/cli"
	"github.com/xiaobinqt/libcompose/project"
)

// ProjectFactory is an interface that helps creating libcompose project.
type ProjectFactory interface {
	// Create creates a libcompose project from the command line options (urfave cli context).
	Create(c *cli.Context) (project.APIProject, error)
}
