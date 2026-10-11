package docker

import (
	"context"
	"openeff/intera/internal/sandbox/backends"
)

type DockerBackend struct {
	Observer
	DockerMouse
	DokerKeyboard
	DockerNetworkTunnel
	DockerTerminal
}

func (b *DockerBackend) Run(ctx context.Context, cfg backends.RunConfig) (err error)
func (b *DockerBackend) Restart(ctx context.Context) (err error)
func (b *DockerBackend) Stop(ctx context.Context) (err error)
func (b *DockerBackend) Close(ctx context.Context) (err error)

type Observer struct {
	DockerScreenshot
	DockerStream
}

var _ backends.BackendSandbox = (*DockerBackend)(nil)
var _ backends.Observer = (*Observer)(nil)
