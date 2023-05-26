package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/xiaobinqt/libcompose/docker"
	"github.com/xiaobinqt/libcompose/docker/ctx"
	"github.com/xiaobinqt/libcompose/project"
	"github.com/xiaobinqt/libcompose/project/options"
)

func main() {
	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{"docker-compose.yml"},
			ProjectName:  "yeah-compose",
		},
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	err = project.Up(context.Background(), options.Up{})

	if err != nil {
		log.Fatal(err)
	}
}
