package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerScreenshot struct {
}

func (s *DockerScreenshot) Screenshot(ctx context.Context) (data []byte, err error)
func (s *DockerScreenshot) ScreenSize(ctx context.Context) (width int, height int, err error)

var _ backends.Screenshot = (*DockerScreenshot)(nil)
