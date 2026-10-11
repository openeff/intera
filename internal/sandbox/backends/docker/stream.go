package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerStream struct {
}

func (s *DockerStream) Start(ctx context.Context) (link string, err error)
func (s *DockerStream) Stop(ctx context.Context) (err error)
func (s *DockerStream) StreamURL() (link string)

var _ backends.Stream = (*DockerStream)(nil)
